package main

import (
	"os"
	"testing"
)

func TestReplaceString(t *testing.T) {
	content := "Hello, ${JAVA_HOME}"
	_ = os.Setenv("JAVA_HOME", "/usr/bin")
	actual := ReplaceString(content)
	expected := "Hello, /usr/bin"
	if actual != expected {
		t.Errorf("Expected : %v, actual is '%v'", expected, actual)
	}
}

func TestReplaceString2(t *testing.T) {
	content := "Hello, %JAVA_HOME%"
	_ = os.Setenv("JAVA_HOME", "/usr/bin")
	actual := ReplaceString(content)
	expected := "Hello, /usr/bin"
	if actual != expected {
		t.Errorf("Expected : %v, actual is '%v'", expected, actual)
	}
}

func TestReplaceString3(t *testing.T) {
	content := "setx /m JAVA_HOME ${TDP_HOME}/${TDP_LIB}/JAVA/${TDP_CURRENT}"
	_ = os.Setenv("JAVA_HOME", "/usr/bin")
	_ = os.Setenv("TDP_HOME", "tdp")
	_ = os.Setenv("TDP_LIB", "lib")
	_ = os.Setenv("TDP_CURRENT", "current")
	actual := ReplaceString(content)
	expected := "setx /m JAVA_HOME tdp/lib/JAVA/current"
	if actual != expected {
		t.Errorf("Expected : %v, actual is '%v'", expected, actual)
	}
}
