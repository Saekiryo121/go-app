package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type appointment struct {
    ID                   int
    NextAppointmentDate string
    ContractedSales      int
    CurrentContractCount int
    CompanyName          string
    CompanyNameKana      string
    Capital              int
    EmployeesCount       int
    AppointmentDepartment string
    ContactPersonName    string
    ContactPersonKana    string
    URL                  string
}

var appointments = []appointment{
    {ID: 1, NextAppointmentDate: "2024-04-01", ContractedSales: 5000000, CurrentContractCount: 80, CompanyName: "株式会社テストデータ", CompanyNameKana: "カブシキガイシャテストデータ", Capital: 12000000, EmployeesCount: 35, AppointmentDepartment: "開発部", ContactPersonName: "鈴木一郎", ContactPersonKana: "スズキイチロウ", URL: "https://www.testdata.co.jp"},
    {ID: 2, NextAppointmentDate: "2024-04-05", ContractedSales: 2000000, CurrentContractCount: 35, CompanyName: "有限会社デモ", CompanyNameKana: "ユウゲンガイシャデモ", Capital: 3000000, EmployeesCount: 10, AppointmentDepartment: "デモ部", ContactPersonName: "田島みゆき", ContactPersonKana: "タジマミユキ", URL: "https://www.demo.co.jp"},
    {ID: 3, NextAppointmentDate: "2024-04-10", ContractedSales: 8000000, CurrentContractCount: 65, CompanyName: "株式会社サンプルデータ", CompanyNameKana: "カブシキガイシャサンプルデータ", Capital: 15000000, EmployeesCount: 50, AppointmentDepartment: "生産部", ContactPersonName: "山口拓也", ContactPersonKana: "ヤマグチタクヤ", URL: "https://www.sampledata.co.jp"},
    {ID: 4, NextAppointmentDate: "2024-04-15", ContractedSales: 3000000, CurrentContractCount: 40, CompanyName: "有限会社モック", CompanyNameKana: "ユウゲンガイシャモック", Capital: 5000000, EmployeesCount: 25, AppointmentDepartment: "モック部", ContactPersonName: "佐々木美咲", ContactPersonKana: "ササキミサキ", URL: "https://www.mock.co.jp"},
    {ID: 5, NextAppointmentDate: "2024-04-20", ContractedSales: 6000000, CurrentContractCount: 55, CompanyName: "株式会社データサンプル", CompanyNameKana: "カブシキガイシャデータサンプル", Capital: 8000000, EmployeesCount: 30, AppointmentDepartment: "営業部", ContactPersonName: "中村健太郎", ContactPersonKana: "ナカムラケンタロウ", URL: "https://www.datasample.co.jp"},
    {ID: 6, NextAppointmentDate: "2024-04-25", ContractedSales: 7000000, CurrentContractCount: 75, CompanyName: "有限会社テストデータ", CompanyNameKana: "ユウゲンガイシャテストデータ", Capital: 10000000, EmployeesCount: 45, AppointmentDepartment: "テスト部", ContactPersonName: "小林奈緒", ContactPersonKana: "コバヤシナオ", URL: "https://www.testdata.co.jp"},
    {ID: 7, NextAppointmentDate: "2024-05-01", ContractedSales: 9000000, CurrentContractCount: 30, CompanyName: "株式会社プロトタイプデータ", CompanyNameKana: "カブシキガイシャプロトタイプデータ", Capital: 18000000, EmployeesCount: 60, AppointmentDepartment: "プロトタイプ部", ContactPersonName: "田中良太", ContactPersonKana: "タナカリョウタ", URL: "https://www.prototypedata.co.jp"},
    {ID: 8, NextAppointmentDate: "2024-05-05", ContractedSales: 4000000, CurrentContractCount: 50, CompanyName: "有限会社イメージデータ", CompanyNameKana: "ユウゲンガイシャイメージデータ", Capital: 12000000, EmployeesCount: 40, AppointmentDepartment: "企画部", ContactPersonName: "岡本みさき", ContactPersonKana: "オカモトミサキ", URL: "https://www.imagedata.co.jp"},
    {ID: 9, NextAppointmentDate: "2024-05-10", ContractedSales: 2500000, CurrentContractCount: 25, CompanyName: "株式会社デザインデータ", CompanyNameKana: "カブシキガイシャデザインデータ", Capital: 7000000, EmployeesCount: 20, AppointmentDepartment: "デザイン部", ContactPersonName: "加藤和也", ContactPersonKana: "カトウカズヤ", URL: "https://www.designdata.co.jp"},
    {ID: 10, NextAppointmentDate: "2024-05-15", ContractedSales: 12000000, CurrentContractCount: 60, CompanyName: "有限会社コンサルトデータ", CompanyNameKana: "ユウゲンガイシャコンサルトデータ", Capital: 25000000, EmployeesCount: 70, AppointmentDepartment: "コンサルティング部", ContactPersonName: "伊藤みさ", ContactPersonKana: "イトウミサ", URL: "https://www.consultdata.co.jp"},
}

func main() {
	corsHandler := func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

					if r.Method == "OPTIONS" {
							return
					}

					h.ServeHTTP(w, r)
			})
	}

	http.Handle("/appointments", corsHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			appointmentsJSON, err := json.Marshal(appointments)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Write(appointmentsJSON)
	})))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
