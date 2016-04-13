package main

import (
	"encoding/json"
	"fmt"
)

type Kube struct {
	Node   Node    `json:"Node"`
	Serv   Service `json:"Service"`
	Checks []Check `json:"Checks"`
}

type Node struct {
	Node string `json:"Node"`
}

type Service struct {
	ID string `json:"ID"`
}

type Check struct {
	CheckID string `json:"CheckID"`
}

func main() {
	k := &[]Kube{}
	if err := json.Unmarshal([]byte(_jsonSample), k); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", k)
}

var _jsonSample = `
[{"Node":{"Node":"kube-node01.dev.ephorus-labs.com","Address":"10.25.120.30","CreateIndex":56573,"ModifyIndex":366063},"Service":{"ID":"nomad-registered-service-21a0fc2f-1d7e-a9e6-a0e4-e87567de276c","Service":"archive-report-pdfgenerator","Tags":["archive","report-pdfgenerator","docker"],"Address":"10.25.120.30","Port":47385,"EnableTagOverride":false,"CreateIndex":366052,"ModifyIndex":366063},"Checks":[{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"0df455ba3ab6c886b284c6f75a4caccd652475a4","Name":"alive","Status":"passing","Notes":"","Output":"TCP connect 10.25.120.30:47385: Success","ServiceID":"nomad-registered-service-21a0fc2f-1d7e-a9e6-a0e4-e87567de276c","ServiceName":"archive-report-pdfgenerator","CreateIndex":366054,"ModifyIndex":366063},{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"Agent alive and reachable","ServiceID":"","ServiceName":"","CreateIndex":56573,"ModifyIndex":56573}]},{"Node":{"Node":"kube-node01.dev.ephorus-labs.com","Address":"10.25.120.30","CreateIndex":56573,"ModifyIndex":366063},"Service":{"ID":"nomad-registered-service-ae3575bd-c28c-38ea-b044-46f77b03486c","Service":"archive-report-pdfgenerator","Tags":["archive","report-pdfgenerator","docker"],"Address":"10.25.120.30","Port":46740,"EnableTagOverride":false,"CreateIndex":366051,"ModifyIndex":366060},"Checks":[{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"d71549a27dd4ff7a546f710b2d2d8ba51a79eb37","Name":"alive","Status":"passing","Notes":"","Output":"TCP connect 10.25.120.30:46740: Success","ServiceID":"nomad-registered-service-ae3575bd-c28c-38ea-b044-46f77b03486c","ServiceName":"archive-report-pdfgenerator","CreateIndex":366053,"ModifyIndex":366060},{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"Agent alive and reachable","ServiceID":"","ServiceName":"","CreateIndex":56573,"ModifyIndex":56573}]},{"Node":{"Node":"kube-node01.dev.ephorus-labs.com","Address":"10.25.120.30","CreateIndex":56573,"ModifyIndex":366063},"Service":{"ID":"nomad-registered-service-e94142f2-753d-1120-e98f-9a36815f604d","Service":"archive-report-pdfgenerator","Tags":["archive","report-pdfgenerator","docker"],"Address":"10.25.120.30","Port":39279,"EnableTagOverride":false,"CreateIndex":366058,"ModifyIndex":366062},"Checks":[{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"ae06ebbbb3eb15bc2be6211f12a88de1c50ba7e7","Name":"alive","Status":"passing","Notes":"","Output":"TCP connect 10.25.120.30:39279: Success","ServiceID":"nomad-registered-service-e94142f2-753d-1120-e98f-9a36815f604d","ServiceName":"archive-report-pdfgenerator","CreateIndex":366059,"ModifyIndex":366062},{"Node":"kube-node01.dev.ephorus-labs.com","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"Agent alive and reachable","ServiceID":"","ServiceName":"","CreateIndex":56573,"ModifyIndex":56573}]},{"Node":{"Node":"kube-node02.dev.ephorus-labs.com","Address":"10.25.120.31","CreateIndex":56732,"ModifyIndex":366044},"Service":{"ID":"nomad-registered-service-3d0e55c2-12cd-82ea-d294-9a90a480a407","Service":"archive-report-pdfgenerator","Tags":["archive","report-pdfgenerator","docker"],"Address":"10.25.120.31","Port":24615,"EnableTagOverride":false,"CreateIndex":366040,"ModifyIndex":366042},"Checks":[{"Node":"kube-node02.dev.ephorus-labs.com","CheckID":"097c6d7800e1eb55a83fc3974ecab2c13f219256","Name":"alive","Status":"passing","Notes":"","Output":"TCP connect 10.25.120.31:24615: Success","ServiceID":"nomad-registered-service-3d0e55c2-12cd-82ea-d294-9a90a480a407","ServiceName":"archive-report-pdfgenerator","CreateIndex":366041,"ModifyIndex":366042},{"Node":"kube-node02.dev.ephorus-labs.com","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"Agent alive and reachable","ServiceID":"","ServiceName":"","CreateIndex":56732,"ModifyIndex":56732}]},{"Node":{"Node":"kube-node02.dev.ephorus-labs.com","Address":"10.25.120.31","CreateIndex":56732,"ModifyIndex":366044},"Service":{"ID":"nomad-registered-service-900f236a-82cc-3574-d165-f66fbb859a36","Service":"archive-report-pdfgenerator","Tags":["archive","report-pdfgenerator","docker"],"Address":"10.25.120.31","Port":22472,"EnableTagOverride":false,"CreateIndex":366038,"ModifyIndex":366044},"Checks":[{"Node":"kube-node02.dev.ephorus-labs.com","CheckID":"6e18708fe3db513cf656b9989fdf9a76f5984888","Name":"alive","Status":"passing","Notes":"","Output":"TCP connect 10.25.120.31:22472: Success","ServiceID":"nomad-registered-service-900f236a-82cc-3574-d165-f66fbb859a36","ServiceName":"archive-report-pdfgenerator","CreateIndex":366039,"ModifyIndex":366044},{"Node":"kube-node02.dev.ephorus-labs.com","CheckID":"serfHealth","Name":"Serf Health Status","Status":"passing","Notes":"","Output":"Agent alive and reachable","ServiceID":"","ServiceName":"","CreateIndex":56732,"ModifyIndex":56732}]}]`
