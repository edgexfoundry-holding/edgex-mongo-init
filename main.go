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

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func main() {
	url := "localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		println("Fatal error during execution: " + err.Error())
		os.Exit(1)
	}

	//db=db.getSiblingDB('admin')
	//db=db.getSiblingDB('authorization')
	db := mgo.Database{
		Session: session,
		Name:    "authorization",
	}

	//db.createUser({ user: "admin",pwd: "password",roles: [ { role: "readWrite", db: "authorization" } ]});
	err = db.UpsertUser(&mgo.User{
		Username: "admin",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	////Create keystore collection
	//db.createCollection("keyStore");
	//keyStore := mgo.Collection{
	//	Database: &db,
	//	Name: "keyStore",
	//	FullName: "db.keyStore",
	//}
	//db.keyStore.insert( { xDellAuthKey: "x-dell-auth-key", secretKey: "EDGEX_SECRET_KEY" } );
	err = db.C("keystore").Insert(bson.D{
		{"xDellAuthKey", "x-dell-auth-key"},
		{"secretKey", "EDGEX_SECRET_KEY"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	////Create Service Mapping
	//db.createCollection("serviceMapping");
	//db.serviceMapping.insert( { serviceName: "coredata", serviceUrl: "http://localhost:48080/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "coredata"},
		{"serviceUrl", "http://localhost:48080/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "metadata", serviceUrl: "http://localhost:48081/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "metadata"},
		{"serviceUrl", "http://localhost:48081/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "command", serviceUrl: "http://localhost:48082/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "command"},
		{"serviceUrl", "http://localhost:48082/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "rules", serviceUrl: "http://localhost:48084/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "rules"},
		{"serviceUrl", "http://localhost:48084/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "notifications", serviceUrl: "http://localhost:48060/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "notifications"},
		{"serviceUrl", "http://localhost:48060/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "logging", serviceUrl: "http://localhost:48061/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "logging"},
		{"serviceUrl", "http://localhost:48061/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.serviceMapping.insert( { serviceName: "export-client", serviceUrl: "http://localhost:48071/" });
	err = db.C("serviceMapping").Insert(bson.D{
		{"serviceName", "export-client"},
		{"serviceUrl", "http://localhost:48071/"}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//
	//db=db.getSiblingDB('admin')
	db = mgo.Database{
		Session: session,
		Name:    "admin",
	}

	//db.system.users.remove({});
	err = db.C("system.users").DropCollection()
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.system.version.remove({});
	err = db.C("system.version").DropCollection()
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.system.version.insert({ "_id" : "authSchema", "currentVersion" : 3 });
	err = db.C("system.users").Insert(bson.D{{
		"_id", "authSchema"},
		{"currentVersion", 3}})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//db=db.getSiblingDB('admin')

	//db.createUser({ user: "admin",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "admin" }
	//]
	//});
	//
	err = db.UpsertUser(&mgo.User{
		Username: "admin",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//db=db.getSiblingDB('metadata')
	db = mgo.Database{
		Session: session,
		Name:    "metadata",
	}
	//db.createUser({ user: "meta",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "metadata" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "meta",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//db.createCollection("addressable");
	addressable := mgo.Collection{
		Database: &db,
		Name:     "addressable",
		FullName: "db.addressable",
	}
	err = addressable.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//TODO unnecessary?
	//db.addressable.createIndex({name: 1}, {unique: true});

	//db.createCollection("command");
	command := mgo.Collection{
		Database: &db,
		Name:     "command",
		FullName: "db.command",
	}
	err = command.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//db.createCollection("device");
	device := mgo.Collection{
		Database: &db,
		Name:     "device",
		FullName: "db.device",
	}
	err = device.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.device.createIndex({name: 1}, {unique: true});

	//db.createCollection("deviceManager");
	deviceManager := mgo.Collection{
		Database: &db,
		Name:     "deviceManager",
		FullName: "db.deviceManager",
	}
	err = deviceManager.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.deviceManager.createIndex({name: 1}, {unique: true});

	//db.createCollection("deviceProfile");
	deviceProfile := mgo.Collection{
		Database: &db,
		Name:     "deviceProfile",
		FullName: "db.deviceProfile",
	}
	err = deviceProfile.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.deviceProfile.createIndex({name: 1}, {unique: true});

	//db.createCollection("deviceReport");
	deviceReport := mgo.Collection{
		Database: &db,
		Name:     "deviceReport",
		FullName: "db.deviceReport",
	}
	err = deviceReport.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.deviceReport.createIndex({name: 1}, {unique: true});

	//db.createCollection("deviceService");
	deviceService := mgo.Collection{
		Database: &db,
		Name:     "deviceService",
		FullName: "db.deviceService",
	}
	err = deviceService.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.deviceService.createIndex({name: 1}, {unique: true});

	//db.createCollection("provisionWatcher");
	provisionWatcher := mgo.Collection{
		Database: &db,
		Name:     "provisionWatcher",
		FullName: "db.provisionWatcher",
	}
	err = provisionWatcher.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.provisionWatcher.createIndex({name: 1}, {unique: true});

	//db.createCollection("schedule");
	schedule := mgo.Collection{
		Database: &db,
		Name:     "schedule",
		FullName: "db.schedule",
	}
	err = schedule.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.schedule.createIndex({name: 1}, {unique: true});

	//db.createCollection("scheduleEvent");
	scheduleEvent := mgo.Collection{
		Database: &db,
		Name:     "scheduleEvent",
		FullName: "db.scheduleEvent",
	}
	err = scheduleEvent.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.scheduleEvent.createIndex({name: 1}, {unique: true});
	//

	//db=db.getSiblingDB('coredata')
	db = mgo.Database{
		Session: session,
		Name:    "coredata",
	}
	//db.createUser({ user: "core",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "coredata" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "core",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("event");
	event := mgo.Collection{
		Database: &db,
		Name:     "event",
		FullName: "db.event",
	}
	err = event.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//TODO what
	//db.event.createIndex({"device": 1}, {unique: false});

	//db.createCollection("reading");
	reading := mgo.Collection{
		Database: &db,
		Name:     "reading",
		FullName: "db.reading",
	}
	err = reading.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	//db.createCollection("valueDescriptor");
	valueDescriptor := mgo.Collection{
		Database: &db,
		Name:     "valueDescriptor",
		FullName: "db.valueDescriptor",
	}
	err = valueDescriptor.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//TODO what
	//db.reading.createIndex({"device": 1}, {unique: false});
	//db.valueDescriptor.createIndex({name: 1}, {unique: true});

	//
	//db=db.getSiblingDB('rules_engine_db')
	db = mgo.Database{
		Session: session,
		Name:    "rules_engine_db",
	}
	//db.createUser({ user: "rules_engine_user",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "rules_engine_db" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "rules_engine_user",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//
	//db=db.getSiblingDB('notifications')
	db = mgo.Database{
		Session: session,
		Name:    "notifications",
	}
	//db.createUser({ user: "notifications",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "notifications" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "notifications",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("notification");
	notification := mgo.Collection{
		Database: &db,
		Name:     "notification",
		FullName: "db.notification",
	}
	err = notification.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("transmission");
	transmission := mgo.Collection{
		Database: &db,
		Name:     "transmission",
		FullName: "db.transmission",
	}
	err = transmission.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("subscription");
	subscription := mgo.Collection{
		Database: &db,
		Name:     "subscription",
		FullName: "db.subscription",
	}
	err = subscription.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.notification.createIndex({slug: 1}, {unique: true});
	//db.subscription.createIndex({slug: 1}, {unique: true});
	//
	//db=db.getSiblingDB('scheduler')
	db = mgo.Database{
		Session: session,
		Name:    "scheduler",
	}
	//db.createUser({ user: "scheduler",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "scheduler" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "scheduler",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("interval");
	interval := mgo.Collection{
		Database: &db,
		Name:     "interval",
		FullName: "db.interval",
	}
	err = interval.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("intervalAction");
	intervalAction := mgo.Collection{
		Database: &db,
		Name:     "intervalAction",
		FullName: "db.intervalAction",
	}
	err = intervalAction.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.interval.createIndex({name: 1}, {unique: true});
	//db.intervalAction.createIndex({name: 1}, {unique: true});
	//
	//db=db.getSiblingDB('logging')
	db = mgo.Database{
		Session: session,
		Name:    "logging",
	}
	//db.createUser({ user: "logging",
	//pwd: "password",
	//roles: [
	//{ role: "readWrite", db: "logging" }
	//]
	//});
	err = db.UpsertUser(&mgo.User{
		Username: "logging",
		Password: "password",
		Roles: []mgo.Role{
			mgo.RoleReadWrite,
		},
	})
	if err != nil {
		println("Error during execution: " + err.Error())
	}
	//db.createCollection("logEntry");
	logEntry := mgo.Collection{
		Database: &db,
		Name:     "logEntry",
		FullName: "db.logEntry",
	}
	err = logEntry.Create(&mgo.CollectionInfo{})
	if err != nil {
		println("Error during execution: " + err.Error())
	}

	session.Close()
}
