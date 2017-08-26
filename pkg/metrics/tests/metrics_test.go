package tests

import (
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/metrics"
	"github.com/grafana/grafana/pkg/metrics/graphite"
	p "github.com/grafana/grafana/pkg/metrics/prometheus"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetricCounters(t *testing.T) {
	Convey("metrics package", t, func() {
		mc := metrics.MetricFactories{}

		Convey("Can use two metrics frameworks", func() {

			graphite.Init(&graphite.GraphiteSettings{
				IntervalSeconds: 10,
			}, mc)
			p.Init(mc)

			metrics.Init(&metrics.MetricSettings{}, mc)
			So(len(mc), ShouldEqual, 2)
		})

		Convey("no configured clients show not cause panic", func() {
			metrics.Init(&metrics.MetricSettings{}, mc)

			So(len(mc), ShouldEqual, 0)

			counter := mc.RegCounter("test counter", "tag", "value")
			counter.Inc(1)
			timer := mc.RegTimer("test timer", "tag", "value")
			duration, _ := time.ParseDuration("1m")
			timer.Update(duration)
			gauge := mc.RegGauge("test gauge", "tag", "value")
			gauge.Update(1)
		})
	})
}
