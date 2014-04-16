// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tGorpController struct {}
var GorpController tGorpController


func (_ tGorpController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Begin", args).Url
}

func (_ tGorpController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Commit", args).Url
}

func (_ tGorpController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Rollback", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tCategory struct {}
var Category tCategory


func (_ tCategory) Index(
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Category.Index", args).Url
}

func (_ tCategory) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Category.New", args).Url
}

func (_ tCategory) Add(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Add", args).Url
}

func (_ tCategory) Edit(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Category.Edit", args).Url
}

func (_ tCategory) Update(
		category interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Category.Update", args).Url
}


type tLanguage struct {}
var Language tLanguage


func (_ tLanguage) Index(
		page int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	return revel.MainRouter.Reverse("Language.Index", args).Url
}

func (_ tLanguage) Edit(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Language.Edit", args).Url
}

func (_ tLanguage) UpdateLanguage(
		language interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "language", language)
	return revel.MainRouter.Reverse("Language.UpdateLanguage", args).Url
}

func (_ tLanguage) New(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Language.New", args).Url
}

func (_ tLanguage) AddLanguage(
		language interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "language", language)
	return revel.MainRouter.Reverse("Language.AddLanguage", args).Url
}

func (_ tLanguage) Delete(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Language.Delete", args).Url
}


