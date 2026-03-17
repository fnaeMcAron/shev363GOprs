// go run main.go -f hosts.txt -monitor -interval 10 -log ping.log
// go run main.go -f hosts.txt -monitor
// go run main.go -f hosts.txt -c 3

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"
)

type PingResult struct {
	Host      string
	Success   bool
	Latency   time.Duration
	Attempt   int
	MaxTrials int
	Timestamp time.Time
	Err       error
}

func pingHost(host string) (bool, time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", host+":80", 2*time.Second)
	latency := time.Since(start)
	if err != nil {
		start = time.Now()
		conn2, err2 := net.DialTimeout("tcp", host+":443", 2*time.Second)
		latency = time.Since(start)
		if err2 != nil {
			return false, 0, err2
		}
		conn2.Close()
		return true, latency, nil
	}
	conn.Close()
	return true, latency, nil
}

func readHosts(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
	}
	var hosts []string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			hosts = append(hosts, line)
		}
	}
	if len(hosts) == 0 {
		return nil, fmt.Errorf("файл %s пуст или не содержит хостов", filename)
	}
	return hosts, nil
}

func openLogFile(name string) (*os.File, error) {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть лог-файл: %w", err)
	}
	return f, nil
}

func formatStatus(success bool) string {
	if success {
		return "OK"
	}
	return "Timeout"
}

func formatLatency(success bool, latency time.Duration) string {
	if !success {
		return "0ms"
	}
	return fmt.Sprintf("%dms", latency.Milliseconds())
}

func runOnce(hosts []string, count int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Хост\tСтатус\tВремя ответа\tПопытка")
	fmt.Fprintln(w, "────\t──────\t────────────\t───────")

	for attempt := 1; attempt <= count; attempt++ {
		for _, host := range hosts {
			ok, latency, _ := pingHost(host)
			fmt.Fprintf(w, "%s\t%s\t%s\t%d/%d\n",
				host,
				formatStatus(ok),
				formatLatency(ok, latency),
				attempt, count,
			)
			if attempt < count {
				time.Sleep(1 * time.Second)
			}
		}
	}
	w.Flush()
}

func worker(ctx context.Context, host string, interval time.Duration, resultCh chan<- PingResult) {
	for {
		ok, latency, err := pingHost(host)
		resultCh <- PingResult{
			Host:      host,
			Success:   ok,
			Latency:   latency,
			Timestamp: time.Now(),
			Err:       err,
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		}
	}
}

func listener(resultCh <-chan PingResult, logFile *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	for res := range resultCh {
		ts := res.Timestamp.Format("2006-01-02 15:04:05")
		status := formatStatus(res.Success)
		latency := formatLatency(res.Success, res.Latency)

		fmt.Printf("%s | %-20s | %-7s | %s\n", ts, res.Host, status, latency)

		if logFile != nil {
			line := fmt.Sprintf("%s | %s | %s | %s\n", ts, res.Host, status, latency)
			logFile.WriteString(line)
		}
	}
}

func runMonitor(hosts []string, interval time.Duration, logPath string) {
	logFile, err := openLogFile(logPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Предупреждение: %v\n", err)
	} else {
		defer logFile.Close()
	}

	resultCh := make(chan PingResult, len(hosts)*2)

	ctx, cancel := context.WithCancel(context.Background())

	var listenerWg sync.WaitGroup
	listenerWg.Add(1)
	go listener(resultCh, logFile, &listenerWg)

	var workerWg sync.WaitGroup
	for _, host := range hosts {
		workerWg.Add(1)
		h := host
		go func() {
			defer workerWg.Done()
			worker(ctx, h, interval, resultCh)
		}()
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Мониторинг запущен. Нажмите Ctrl+C для завершения.")
	fmt.Printf("%-19s | %-20s | %-7s | %s\n", "Время", "Хост", "Статус", "Задержка")
	fmt.Println(strings.Repeat("─", 60))

	<-sigChan
	fmt.Println("\nПолучен сигнал завершения, останавливаем воркеры...")

	cancel()
	workerWg.Wait()
	close(resultCh)
	listenerWg.Wait()

	fmt.Println("Мониторинг завершён.")
}

func main() {
	hostsFile := flag.String("f", "hosts.txt", "Файл со списком хостов")
	count := flag.Int("c", 1, "Количество повторов проверки (режим разовой проверки)")
	monitor := flag.Bool("monitor", false, "Режим постоянного мониторинга")
	interval := flag.Int("interval", 5, "Интервал проверки в секундах (только для -monitor)")
	logPath := flag.String("log", "monitor.log", "Путь к файлу лога (только для -monitor)")
	flag.Parse()

	hosts, err := readHosts(*hostsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	if *monitor {
		runMonitor(hosts, time.Duration(*interval)*time.Second, *logPath)
	} else {
		runOnce(hosts, *count)
	}
}
