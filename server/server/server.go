package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	listener net.Listener   // TCP слушатель
	wg       sync.WaitGroup // Группа ожидания для завершения всех соединений
	stop     chan struct{}  // Канал для сигнала остановки сервера
}

func NewServer() *Server {
	return &Server{
		stop: make(chan struct{}), // Инициализация канала остановки
	}
}

func (s *Server) Start(address, wisdoms string) error {
	// Load quotes
	if err := LoadQuotes(wisdoms); err != nil {
		return fmt.Errorf("failed to load quotes: %w", err)
	}

	// Open TCP port
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.listener = ln
	log.Printf("Server listening on %s", address)

	// Handle Ctrl+C / SIGTERM
	go s.handleSignals()

	//🔁 Этот цикл:
	//  Постоянно ждёт новые подключения.
	//	Обрабатывает каждое соединение в отдельной горутине.
	//	Учитывает активные соединения с помощью sync.WaitGroup.
	//	Адекватно реагирует на остановку по сигналу (Ctrl+C, SIGTERM), прерывая цикл и корректно завершая работу.
	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-s.stop:
				log.Println("Server is shutting down.")
				return nil
			default:
				log.Printf("Accept error: %v", err)
				continue
			}
		}
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			handleConnection(conn)
		}()
	}
}

func (s *Server) handleSignals() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	s.Stop()
}

func (s *Server) Stop() {
	log.Println("Stopping server...")
	close(s.stop)
	if s.listener != nil {
		s.listener.Close()
	}
	s.wg.Wait()
	log.Println("Server shutdown complete.")
}
