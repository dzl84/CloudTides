package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"strconv"
	"tides-server/pkg/config"
	"tides-server/pkg/models"
	"tides-server/pkg/restapi/operations/org"
	"time"
)


func AddOrgHandler(params org.AddOrgParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return org.NewAddOrgUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	body := params.ReqBody
	db := config.GetDB()
	var orgOld models.Org;
	if db.Where("org_name = ?", body.Name).First(&orgOld).RowsAffected == 1 {
		//if new org name is the same as old one, just forbid it
		return org.NewAddOrgForbidden().WithPayload(&org.AddOrgForbiddenBody{
			Message: "Insert DB error:, org name is already existed",
		})
	}
	if db.Unscoped().Where("org_name = ?", body.Name).First(&orgOld).RowsAffected == 1 {
		// if repeat with deleted org, just delete it
		db.Unscoped().Delete(&orgOld)
	}
	newOrg := models.Org{
		OrgName: body.Name,
	}
	if err := db.Create(&newOrg).Error; err != nil {
		return org.NewAddOrgUnauthorized().WithPayload(&org.AddOrgUnauthorizedBody{
			Message: "Insert DB error:" + err.Error(),
		})
	}

	newLog := models.Log{
		UserID: uid,
		Operation: "Add Org:" + body.Name,
		Time: time.Now(),
		Status: "Succeed",
	}
	if err := db.Create(&newLog).Error; err != nil {
		return org.NewAddOrgForbidden().WithPayload(&org.AddOrgForbiddenBody{
			Message: "Insert Log DB error:" + err.Error(),
		})
	}

	return org.NewAddOrgOK().WithPayload(&org.AddOrgOKBody{
		Message: "succeed",
	})
}

func ListOrgHandler(params org.ListOrgParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return org.NewListOrgUnauthorized()
	}
	var orgs []*models.Org
	db := config.GetDB()
	db.Find(&orgs)
	var reponse []*org.ListOrgOKBodyItems0
	for _, tmpOrg := range orgs {
		// get resource id  from resource new table
		var resources []*models.ResourceNew
		db.Where("org_id = ? ", tmpOrg.ID).Order("id").Find(&resources)
		var totalCpu, totalRAM, totalDisk, curCpu, curRAM, curDisk float64
		for _, tmpRes :=  range resources{
			var resUsages models.ResourceUsage
			db.Where("id = ? ", tmpRes.ID).Find(&resUsages)
			curCpu += resUsages.CurrentCPU
			curRAM += resUsages.CurrentRAM
			curDisk += resUsages.CurrentDisk
			totalCpu += resUsages.TotalCPU
			totalRAM += resUsages.TotalRAM
			totalDisk += resUsages.TotalDisk
		}
		newResult := org.ListOrgOKBodyItems0{
			ID: int64(tmpOrg.ID),
			Name: tmpOrg.OrgName,
			CurrentCPU: curCpu,
			CurrentRAM: curRAM,
			CurrentDisk: curDisk,
			TotalCPU:  totalCpu,
			TotalRAM:  totalRAM,
			TotalDisk:  totalDisk,
		}

		reponse = append(reponse, &newResult)
	}
	return org.NewListOrgOK().WithPayload(reponse)
}

func DeleteOrgHandler(params org.DeleteOrgParams) middleware.Responder {
	if !VerifyUser(params.HTTPRequest) {
		return org.NewDeleteOrgUnauthorized()
	}
	uid, _ := ParseUserIDFromToken(params.HTTPRequest)
	db := config.GetDB()
	var pol models.Org
	db.Where("id = ? ", params.ID).Find(&pol)
	var orgName = pol.OrgName
	if db.Where("id = ? ", params.ID).Delete(&pol).RowsAffected == 0 {
		return org.NewDeleteOrgNotFound().WithPayload(&org.DeleteOrgNotFoundBody{
			Message: "org with id " + strconv.FormatInt(params.ID, 10) + " is not found in database",
		})
	}

	//Delete User in that org
	var users []*models.User
	db.Where("org_name = ? ", orgName).Delete(&users)

	newLog := models.Log{
		UserID: uid,
		Operation: "Delete Org:" + orgName,
		Time: time.Now(),
		Status: "Succeed",
	}
	if db.Create(&newLog).Error != nil {
		return org.NewAddOrgForbidden()
	}

	return org.NewDeleteOrgOK().WithPayload(&org.DeleteOrgOKBody{
		Message: "success",
	})
}

