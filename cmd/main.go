package main

import (
 "fmt"
 "log"

 "github.com/imamirhossein43/v2ray_user" 

)

func main() {
 serverURL := "https://example.com" // آدرس سرور V2Ray خود را قرار دهید
 token := "your_token" // توکن معتبر خود را وارد کنید

 // ایجاد نمونه V2RayUserManager
 userManager := v2ray_user.NewV2RayUserManager(serverURL, token)

 // اضافه کردن کاربر
 newUser := v2ray_user.User{
  ID:       "user1",
  Email:    "user1@example.com",
  Password: "password1",
 }
 err := userManager.CreateUser(newUser)
 if err != nil {
  log.Fatal(err)
 }

 // حذف کاربر
 err = userManager.DeleteUser("user1")
 if err != nil {
  log.Fatal(err)
 }

 // دریافت وضعیت کاربر
 status, err := userManager.GetUserStatus("user1")
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("وضعیت کاربر: ", status)
}