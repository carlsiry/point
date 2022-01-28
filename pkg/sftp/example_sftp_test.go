package sftp_test

import (
	"github.com/carlsiry/point/pkg/sftp"
	"net/http"
)

func ExamplePoint_Upload() {
	// 1. init sftp point
	point, err := sftp.NewPoint("root", "pwd123456", "192.168.11.1", 22)
	if err != nil {
		// handle err
	}

	// 2. inject inner layer
	uc := UseCaseLayer(point)

	// 3. serve controller
	Serve(Controller(uc))
}

type UseCase func(from string, to string) error

func UseCaseLayer(point sftp.IPoint) UseCase {
	return func(from, to string) error {
		// ... handle sth.
		return point.Upload(from, to)
	}
}

func Controller(useCase UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		from := request.URL.Query().Get("fromPath")
		to := request.URL.Query().Get("toPath")
		if from == "" || to == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := useCase(from, to); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Serve(handler http.HandlerFunc) {
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
