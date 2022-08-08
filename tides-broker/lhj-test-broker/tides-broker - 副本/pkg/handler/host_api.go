package handler

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"

	"tides-broker/pkg/config"
	"tides-broker/pkg/models"
	"tides-broker/pkg/restapi/operations/hosts_api"
)

// api for adding one host
func AddOneHost(params hosts_api.AddOneHostParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return NewAddOneHostUnauthorized()
	// }

	db := config.GetDB()

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.Body

	hostname := body.Hostname
	// search for hostname in the db and check if exists
	checkrst := db.Exec(fmt.Sprintf("SELECT * FROM hosts WHERE hostname = %s;", hostname))
	

	var rec = make(map[string]interface{})
    if checkrst.Find(rec); len(rec) > 0 {
        return hosts_api.NewAddOneHostBadRequest().WithPayload(&hosts_api.AddOneHostBadRequestBody{
			Message: "Hostname already exists",
		})
    }

	newHost := models.Host{
		Hostname: body.Hostname,
    	Policy: body.Policy,
    	Datacenter: body.Datacenter,
		Cluster: body.Cluster,
		IP: body.IP,
    	Port: body.Port,
		Sshkey: body.Sshkey,
    	Username: body.Username,
		Password: body.Password,
	}

	db.Create(&newHost)

	return hosts_api.NewAddOneHostCreated().WithPayload(&hosts_api.AddOneHostCreatedBody{
		Message: "success",
	})
}

// api for listing all the hosts
func GetHosts(params hosts_api.GetHostsParams) middleware.Responder {
	// if !VerifyUser(params.HTTPRequest) {
	// 	return resource.NewListVsphereResourceUnauthorized()
	// }

	// uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()

	// if VerifyAdmin(params.HTTPRequest) {
	// 	db.Find(&resources)
	// } else {
	// 	db.Where("user_id = ?", uid).Find(&resources)
	// }
	var hosts []*models.Host
	result := db.Find(&hosts)

	if result.Error  != nil {
		return hosts_api.NewGetHostsBadRequest().WithPayload(&hosts_api.GetHostsBadRequestBody{
			Message: "Cannot query the results",
		})
	}

	var response []*hosts_api.GetHostsOKBodyResultsItems0
	for _, tmpHost := range hosts {
		newResult := hosts_api.GetHostsOKBodyResultsItems0{
			Hostname: tmpHost.Hostname,
    	    Policy: tmpHost.Policy,
    	    Datacenter: tmpHost.Datacenter,
		    Cluster: tmpHost.Cluster,
		    IP: tmpHost.IP,
    	    Port: tmpHost.Port,
		    Sshkey: tmpHost.Sshkey,
		}

		response = append(response, &newResult)
	}

	return hosts_api.NewGetHostsOK().WithPayload(&hosts_api.GetHostsOKBody{
		Message: "success",
		Results: response,
	})
}


// api for deleting one host
func DeleteHost(params hosts_api.DeleteHostParams) middleware.Responder {
    // if !VerifyUser(params.HTTPRequest) {
    //  return NewAddOneHostUnauthorized()
    // }

    db := config.GetDB()

    // uid, _ := ParseUserIDFromToken(params.HTTPRequest)
    body := params.Body

    hostToDelete := body.Hostname

    rst := db.Exec(`DELETE FROM hosts WHERE hostname = $1;`, hostToDelete)
    if rst.Error != nil {
        return hosts_api.NewDeleteHostBadRequest().WithPayload(&hosts_api.DeleteHostBadRequestBody{
            Message: "Cannot delete host",
        })
    }

    return hosts_api.NewDeleteHostOK().WithPayload(&hosts_api.DeleteHostOKBody{
        Hostname: hostToDelete,
    })
}
