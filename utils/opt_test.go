package utils

import (
	"testing"

	"github.com/streadway/amqp"
	"github.com/tiago4orion/wabbit"
)

func TestConvertOptDefaults(t *testing.T) {
	opt, err := ConvertOpt(nil)

	if err != nil {
		t.Error(err)
		return
	}

	if opt.ContentType != "text/plain" {
		t.Error("Invalid opt content type: %s", opt.ContentType)
	}

	if opt.ContentEncoding != "" {
		t.Error("Invalid opt encoding: %s", opt.ContentEncoding)
	}

	if opt.DeliveryMode != amqp.Transient {
		t.Error("Invalid default delivery mode: %d\n", opt.DeliveryMode)
	}

	if opt.Priority != uint8(0) {
		t.Error("Invalid default priority: %d\n", opt.Priority)
	}

	if len(opt.Headers) != 0 {
		t.Errorf("Invalid value for headers: %v", opt.Headers)
	}
}

func TestConvertOpt(t *testing.T) {
	opt, err := ConvertOpt(wabbit.Option{
		"contentType": "binary/fuzz",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if opt.ContentType != "binary/fuzz" {
		t.Errorf("Wrong value for content type: %s", opt.ContentType)
	}

	opt, err = ConvertOpt(wabbit.Option{
		"contentEncoding": "bleh",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if opt.ContentEncoding != "bleh" {
		t.Errorf("Invalid value for contentEncoding: %s", opt.ContentEncoding)
	}

	opt, err = ConvertOpt(wabbit.Option{
		"contentEncoding": "bleh",
		"contentType":     "binary/fuzz",
	})

	if err != nil {
		t.Error(err)
		return
	}

	if opt.ContentType != "binary/fuzz" {
		t.Errorf("Wrong value for content type: %s", opt.ContentType)
	}

	if opt.ContentEncoding != "bleh" {
		t.Errorf("Invalid value for contentEncoding: %s", opt.ContentEncoding)
	}

	// setting invalid value

	opt, err = ConvertOpt(wabbit.Option{
		"NotExists": "bleh",
	})

	if err == nil {
		t.Errorf("Shall fail...")
		return
	}
}