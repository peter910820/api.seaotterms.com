# api.seaotterms.com

* 因為近期專案都準備前後端完全分離，所以開了一個api專案去集中API管理，減少伺服器需要開的PORT(專案)數量。
* 目前旗下有管理三個專案的API:
    * blog/seaotterms.com: 主站(前後端分離中)
    * gal/gal.seaotterms.com: galgame文章資源分享站(開發中)
    * teach/teach.seaotterms.com: 教學文章站(關閉中)

## 專案架構

### api模組
* 存放API方法，用**站台別**分子模組
### router模組
* 存放站台路由，用**站台別**分子模組  
### model模組
* 存放資料表結構，用**資料庫別**分子模組
### dto模組
* 存放DTO結構，用**站台別**分子模組  

>[!NOTICE]
>待補