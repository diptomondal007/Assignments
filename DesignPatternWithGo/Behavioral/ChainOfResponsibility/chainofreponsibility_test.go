package ChainOfResponsibility

import (
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T){
	myWriter := myTestWriter{}
	writerLogger := WriterLogger{
		Writer:    &myWriter,
	}
	secondLogger := SecondLogger{NextChain: &writerLogger}
	chain := FirstLogger{NextChain: &secondLogger}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+ "the word 'hello', third writes to some variable if second found 'hello'", func(t *testing.T) {
		chain.Next("message that breaks the chain\n")
		if myWriter.receivedMessage != nil {
			t.Error("Last link should not receive any message")
		}
		chain.Next("Hello\n")
		if myWriter.receivedMessage == nil || !strings.Contains(*myWriter.receivedMessage, "Hello") {
			t.Fatal("Last link didn't received expected message") }
	})
}
