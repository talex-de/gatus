package core

import "testing"

func TestNewServiceStatus(t *testing.T) {
	service := &Service{Name: "name", Group: "group"}
	serviceStatus := NewServiceStatus(service)
	if serviceStatus.Name != service.Name {
		t.Errorf("expected %s, got %s", service.Name, serviceStatus.Name)
	}
	if serviceStatus.Group != service.Group {
		t.Errorf("expected %s, got %s", service.Group, serviceStatus.Group)
	}
}

func TestServiceStatus_AddResult(t *testing.T) {
	service := &Service{Name: "name", Group: "group"}
	serviceStatus := NewServiceStatus(service)
	for i := 0; i < 50; i++ {
		serviceStatus.AddResult(&Result{})
	}
	if len(serviceStatus.Results) != 20 {
		t.Errorf("expected serviceStatus.Results to not exceed a length of 20")
	}
}
