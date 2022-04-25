# MoodMoment-server  

## 说明

> 由于该小程序会包含多个服务应用，所以采用微服务框架【gozero】

### 一、AdMonitor-api服务  

* datasource
  
  ```bash
  Admonitor:imih7kNDasGJ6dFP@tcp(150.109.61.90:3306)/admonitor?charset=utf8mb4&parseTime=True&loc=Local
  ```

* api生成

  ```bash
  goctl api go -api admonitor.api -dir="../api" --style goZero
  ```

* model生成
  
  ```bash
  goctl model mysql datasource -url="Admonitor:imih7kNDasGJ6dFP@tcp(150.109.61.90:3306)/admonitor" -table="*" -dir="." --style goZero
  ```

* swagger生成

  ```bash
  goctl api plugin -plugin goctl-swagger="swagger -filename admonitor.json -host wx.api.catli.net -basepath /" -api admonitor.api -dir .
  ```
