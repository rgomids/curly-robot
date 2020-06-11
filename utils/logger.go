package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

var (
	logTemplate string = `[%s] [CONTEXT: {%s} | SESSION: {%s}]	(%s)`
)

// Logger contem toda a interface para logs do sistema
type Logger struct {
	context string
	session string
	logFile *os.File
	*log.Logger
}

// NewLogger inicia um novo contexto de log para o sistema
func NewLogger() *Logger {
	f, err := os.OpenFile(fmt.Sprintf("/var/log/%s", os.Getenv("TURBO_PARAKEET_DEFAULT_NAME")), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger := log.Logger{}
	logger.SetOutput(f)

	newLogger := &Logger{
		context: uuid.New().String(),
		logFile: f,
	}

	newLogger.Logger = &logger

	return newLogger
}

// NewSession inicia uma nova sessao para os logs
func (l *Logger) NewSession() *Logger {
	return &Logger{
		context: l.context,
		session: uuid.New().String(),
	}
}

// Info interface de log padrao do sistema
func (l *Logger) Info(text string) {
	log.Println(l.loadTemplate(`INFO`, text))
}

// Error interface de log de erros do sistema
func (l *Logger) Error(err error) {
	log.Println(l.loadTemplate(`ERRO`, err.Error()))
}

// Fatal interface de log fatais do sistema
func (l *Logger) Fatal(err error) {
	log.Fatalln(l.loadTemplate(`FATAL`, err.Error()))
}

func (l *Logger) loadTemplate(level, text string) string {
	return fmt.Sprintf(logTemplate, level, l.context, l.session, text)
}

// GetServerContextID retorna o ID do contexto atual
func (l *Logger) GetServerContextID() string {
	return l.context
}

// CloseLogFile fecha o arquivo de log
func (l *Logger) CloseLogFile() {
	l.logFile.Close()
}
