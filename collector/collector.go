package collector

import (
	"golang.org/x/exp/constraints"
)

type MetricValue interface {
	constraints.Integer | constraints.Float
}
