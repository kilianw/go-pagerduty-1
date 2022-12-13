package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestAutomationActionsScriptActionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"action":{"action_data_reference":{"script":"java --version","invocation_command":"sh"},"action_type":"script","action_classification":"diagnostic","creation_time":"2022-12-12T18:51:42.048162Z","id":"01DF4OBNYKW84FS9CCYVYS1MOS","name":"Script Action created by TF","type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Get("01DF4OBNYKW84FS9CCYVYS1MOS")
	if err != nil {
		t.Fatal(err)
	}

	script := "java --version"
	invocation_command := "sh"
	classification := "diagnostic"
	adf := AutomationActionsActionDataReference{
		Script:            &script,
		InvocationCommand: &invocation_command,
	}
	want := &AutomationActionsAction{
		ID:                   "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:                 "Script Action created by TF",
		CreationTime:         "2022-12-12T18:51:42.048162Z",
		ActionType:           "script",
		Type:                 "action",
		ActionClassification: &classification,
		ActionDataReference:  adf,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsProcessAutomationActionGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{"action":{"action_data_reference":{"process_automation_job_id":"1519578e-a22a-4340-b58f-08194691e10b"},"action_type":"process_automation","creation_time":"2022-12-12T18:51:42.048162Z","id":"01DF4OBNYKW84FS9CCYVYS1MOS","name":"Action created by TF","privileges":{"permissions":["read"]},"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Get("01DF4OBNYKW84FS9CCYVYS1MOS")
	if err != nil {
		t.Fatal(err)
	}

	job_id := "1519578e-a22a-4340-b58f-08194691e10b"
	adf := AutomationActionsActionDataReference{
		ProcessAutomationJobId: &job_id,
	}
	permissions_read := "read"
	want := &AutomationActionsAction{
		ID:                  "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:                "Action created by TF",
		CreationTime:        "2022-12-12T18:51:42.048162Z",
		ActionType:          "process_automation",
		Type:                "action",
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionCreate(t *testing.T) {
	setup()
	defer teardown()

	description := "Description of Action created by TF"
	runner_id := "01DF4O9T1MDPYOUT7SUX9EXZ4R"
	adf_arg := "-arg 123"
	job_id := "1519578e-a22a-4340-b58f-08194691e10b"
	adf := AutomationActionsActionDataReference{
		ProcessAutomationJobId:        &job_id,
		ProcessAutomationJobArguments: &adf_arg,
	}
	input := &AutomationActionsAction{
		Name:                "Action created by TF",
		Description:         &description,
		ActionType:          "process_automation",
		RunnerID:            &runner_id,
		ActionDataReference: adf,
	}

	mux.HandleFunc("/automation_actions/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(AutomationActionsActionPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Action, input) {
			t.Errorf("Request body = %+v, want %+v", v.Action, input)
		}
		w.Write([]byte(`{"action":{"action_data_reference":{"process_automation_job_id":"1519578e-a22a-4340-b58f-08194691e10b","process_automation_job_arguments":"-arg 123"},"action_type":"process_automation","creation_time":"2022-12-12T18:51:42.048162Z","description":"Description of Action created by TF","id":"01DF4OBNYKW84FS9CCYVYS1MOS","last_run":"2022-12-12T18:52:11.937747Z","last_run_by":{"id":"PINL781","type":"user_reference"},"modify_time":"2022-12-12T18:51:42.048162Z","name":"Action created by TF","privileges":{"permissions":["read"]},"runner":"01DF4O9T1MDPYOUT7SUX9EXZ4R","runner_type":"runbook","services":[{"id":"PQWQ0U6","type":"service_reference"}],"teams":[{"id":"PZ31N6S","type":"team_reference"}],"type":"action"}}`))
	})

	resp, _, err := client.AutomationActionsAction.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	runner_type_runbook := "runbook"
	modify_time := "2022-12-12T18:51:42.048162Z"
	permissions_read := "read"
	want := &AutomationActionsAction{
		ID:           "01DF4OBNYKW84FS9CCYVYS1MOS",
		Name:         "Action created by TF",
		Description:  &description,
		CreationTime: "2022-12-12T18:51:42.048162Z",
		ActionType:   "process_automation",
		Type:         "action",
		RunnerID:     &runner_id,
		RunnerType:   &runner_type_runbook,
		Teams: []*TeamReference{
			{
				Type: "team_reference",
				ID:   "PZ31N6S",
			},
		},
		Services: []*ServiceReference{
			{
				Type: "service_reference",
				ID:   "PQWQ0U6",
			},
		},
		ActionDataReference: adf,
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
		ModifyTime: &modify_time,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsActionDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/actions/01DF4OBNYKW84FS9CCYVYS1MOS", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.AutomationActionsAction.Delete("01DF4OBNYKW84FS9CCYVYS1MOS"); err != nil {
		t.Fatal(err)
	}
}
