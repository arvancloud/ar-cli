package utl

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	apiErrors "github.com/masihyeganeh/ar-cli/pkg/api/errors"
	"github.com/masihyeganeh/ar-cli/pkg/api/responses"
	"io"
	"log"
	"os"
	"strings"
)

const (
	DefaultErrorExitCode = 1
)

// #TODO improve check error to print better

// fatal prints the message (if provided) and then exits.
func fatalErrHandler(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}
		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code when an api error happens.
func CheckApiErr(err error) {
	if err != nil {
		errMsg := err.Error()
		if err, ok := err.(apiErrors.GenericApiCallError); ok {
			if err, ok := err.Model().(responses.InlineResponse422); ok {
				errMsg += "\n\n" + err.Message
				for field, errorDetails := range err.Errors {
					errMsg += fmt.Sprintf("\n%s: %s", field, strings.Join(errorDetails, ", "))
				}
			}
		}
		CheckErr(errors.New(errMsg))
	}
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code. Unrecognized errors will be printed with an "error: " prefix.
func CheckErr(err error) {
	checkErr(err, fatalErrHandler)
}

// checkErr formats a given error as a string and calls the passed handleErr
func checkErr(err error, handleErr func(string, int)) {
	if err == nil {
		return
	}
	fmt.Println(err)
	handleErr("", DefaultErrorExitCode)
}

// ReadInput prints explain and repeat printing inputExplain to out and reads a string from in.
//   If input is empty and defaultVal is set returns default value
//   If defaultVal is not set, tries to validate input using validate
func ReadInput(inputExplain, defaultVal string, out io.Writer, in io.Reader, validate func(string) (bool, error)) string {
	reader := bufio.NewReader(in)
	for {
		_, err := fmt.Fprint(out, inputExplain)
		if err != nil {
			log.Println(err)
		}
		i, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintf(out, "Error: %s\n", err.Error())
			if err != nil {
				log.Println(err)
			}
		} else {
			i = strings.TrimSpace(i)
			if len(i) == 0 && len(defaultVal) > 0 {
				return defaultVal
			}
			valid, err := validate(i)
			if valid {
				return i
			}
			_, err = fmt.Fprintf(out, "Error: %s\n", err.Error())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// EncodeInner tries to encode one of non-empty inner parts of a struct as a workaround for stupid Go bug
func EncodeInner(inners ...interface{}) ([]byte, error) {
	var result []byte
	var err error
	for _, inner := range inners {
		if result, err = json.Marshal(inner); err == nil && len(result) > 2 {
			return result, err
		}
	}
	return []byte("{}"), nil
}
