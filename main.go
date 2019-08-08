/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

// main package contains the entry point for the application and the starting business logic.
package main

import (
	"os"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func main() {
	LoggingClient := logger.NewClient("mongo-init", false, "mongo-init.log", models.DebugLog)

	url := "localhost:27017"

	session, err := mgo.Dial(url)
	if err != nil {
		LoggingClient.Error("Fatal error during execution: " + err.Error())
		os.Exit(1)
	}

	defer session.Close()

	db := mgo.Database{
		Session: session,
		Name:    "authorization",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "admin",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("keystore").Insert(bson.D{
		{Name: "xDellAuthKey", Value: "x-dell-auth-key"},
		{Name: "secretKey", Value: "EDGEX_SECRET_KEY"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "coredata"},
		{Name: "serviceUrl", Value: "http://localhost:48080/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "metadata"},
		{Name: "serviceUrl", Value: "http://localhost:48081/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "command"},
		{Name: "serviceUrl", Value: "http://localhost:48082/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "rules"},
		{Name: "serviceUrl", Value: "http://localhost:48084/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "notifications"},
		{Name: "serviceUrl", Value: "http://localhost:48060/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "logging"},
		{Name: "serviceUrl", Value: "http://localhost:48061/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
	err = db.C("serviceMapping").Insert(bson.D{
		{Name: "serviceName", Value: "export-client"},
		{Name: "serviceUrl", Value: "http://localhost:48071/"}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "admin",
	}

	_, err = db.C("system.users").RemoveAll(nil)
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	_, err = db.C("system.version").Upsert(nil, bson.D{
		{Name: "_id", Value: "authSchema"},
		{Name: "currentVersion", Value: 3}})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	err = db.UpsertUser(&mgo.User{
		Username: "admin",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "metadata",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "meta",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.addressable.createIndex({name: 1}, {unique: true});
	err = db.C("addressable").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	command := mgo.Collection{
		Database: &db,
		Name:     "command",
		FullName: "db.command",
	}
	err = command.Create(&mgo.CollectionInfo{})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.device.createIndex({name: 1}, {unique: true});
	err = db.C("device").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.deviceManager.createIndex({name: 1}, {unique: true});
	err = db.C("deviceManager").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.deviceProfile.createIndex({name: 1}, {unique: true});
	err = db.C("deviceProfile").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.deviceReport.createIndex({name: 1}, {unique: true});
	err = db.C("deviceReport").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.deviceService.createIndex({name: 1}, {unique: true});
	err = db.C("deviceService").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.provisionWatcher.createIndex({name: 1}, {unique: true});
	err = db.C("provisionWatcher").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.schedule.createIndex({name: 1}, {unique: true});
	err = db.C("schedule").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.scheduleEvent.createIndex({name: 1}, {unique: true});
	err = db.C("scheduleEvent").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "coredata",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "core",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.event.createIndex({"device": 1}, {unique: false});
	err = db.C("event").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.reading.createIndex({"device": 1}, {unique: false});
	err = db.C("reading").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.valueDescriptor.createIndex({name: 1}, {unique: true});
	err = db.C("valueDescriptor").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "rules_engine_db",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "rules_engine_user",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "notifications",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "notifications",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.notification.createIndex({slug: 1}, {unique: true});
	err = db.C("notification").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	transmission := mgo.Collection{
		Database: &db,
		Name:     "transmission",
		FullName: "db.transmission",
	}
	err = transmission.Create(&mgo.CollectionInfo{})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	//db.subscription.createIndex({slug: 1}, {unique: true});
	err = db.C("subscription").EnsureIndex(mgo.Index{Key: []string{"name"}, Name: "name_1", Unique: true})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "scheduler",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "scheduler",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	interval := mgo.Collection{
		Database: &db,
		Name:     "interval",
		FullName: "db.interval",
	}
	err = interval.Create(&mgo.CollectionInfo{})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	intervalAction := mgo.Collection{
		Database: &db,
		Name:     "intervalAction",
		FullName: "db.intervalAction",
	}
	err = intervalAction.Create(&mgo.CollectionInfo{})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	db = mgo.Database{
		Session: session,
		Name:    "logging",
	}

	err = db.UpsertUser(&mgo.User{
		Username: "logging",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}

	logEntry := mgo.Collection{
		Database: &db,
		Name:     "logEntry",
		FullName: "db.logEntry",
	}
	err = logEntry.Create(&mgo.CollectionInfo{})
	if err != nil {
		LoggingClient.Error("Error during execution: " + err.Error())
	}
}
