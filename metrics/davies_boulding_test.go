package metrics_test

import (
	"test/metrics"
	"testing"
)

func TestDaviesBouldinIndex(t *testing.T) {
	d := metrics.DaviesBouldinIndex(X, centroids, labels)
	t.Log(d)
}
