package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/Knetic/govaluate"
)

type Request struct {
	Expr string `json:"expr"`
	Deg  bool   `json:"deg"` // true = graus, false = radianos
}

type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func toFloat(v interface{}) float64 {
	switch v := v.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		return 0
	}
}

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			
		}

		functions := map[string]govaluate.ExpressionFunction{
			"sin": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				if req.Deg {
					x = x * math.Pi / 180
				}
				return math.Sin(x), nil
			},
			"cos": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				if req.Deg {
					x = x * math.Pi / 180
				}
				return math.Cos(x), nil
			},
			"tan": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				if req.Deg {
					x = x * math.Pi / 180
				}
				return math.Tan(x), nil
			},
			"asin": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				val := math.Asin(x)
				if req.Deg {
					val = val * 180 / math.Pi
				}
				return val, nil
			},
			"acos": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				val := math.Acos(x)
				if req.Deg {
					val = val * 180 / math.Pi
				}
				return val, nil
			},
			"atan": func(args ...interface{}) (interface{}, error) {
				x := toFloat(args[0])
				val := math.Atan(x)
				if req.Deg {
					val = val * 180 / math.Pi
				}
				return val, nil
			},
			"ln": func(args ...interface{}) (interface{}, error) {
				return math.Log(toFloat(args[0])), nil
			},
			"log": func(args ...interface{}) (interface{}, error) {
				return math.Log10(toFloat(args[0])), nil
			},
			"sqrt": func(args ...interface{}) (interface{}, error) {
				return math.Sqrt(toFloat(args[0])), nil
			},
			"abs": func(args ...interface{}) (interface{}, error) {
				return math.Abs(toFloat(args[0])), nil
			},
			"floor": func(args ...interface{}) (interface{}, error) {
				return math.Floor(toFloat(args[0])), nil
			},
			"ceil": func(args ...interface{}) (interface{}, error) {
				return math.Ceil(toFloat(args[0])), nil
			},
			"exp": func(args ...interface{}) (interface{}, error) {
				return math.Exp(toFloat(args[0])), nil
			},
			"pi": func(args ...interface{}) (interface{}, error) {
				return math.Pi, nil
			},
			"e": func(args ...interface{}) (interface{}, error) {
				return math.E, nil
			},
		}

		expr, err := govaluate.NewEvaluableExpressionWithFunctions(req.Expr, functions)
		var res Response
		if err != nil {
			res.Error = err.Error()
		} else {
			val, err := expr.Evaluate(nil)
			if err != nil {
				res.Error = err.Error()
			} else {
				res.Result = toFloat(val)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
// linhas de comentário apenas para o card do repositório mostrar como GoLang
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//