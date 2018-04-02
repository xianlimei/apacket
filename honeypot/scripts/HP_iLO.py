#!/usr/bin/env python
# -*- coding: UTF-8 -*-

createUserResp = {
    "@odata.context": "/redfish/v1/$metadata#AccountService/Accounts/Members/$entity", 
    "@odata.id": "/redfish/v1/AccountService/Accounts/3/", 
    "@odata.type": "#ManagerAccount.1.0.0.ManagerAccount", 
    "Description": "iLO User Account", 
    "Id": "3", 
    "Name": "User Account", 
    "Oem": {
        "Hp": {
            "@odata.type": "#HpiLOAccount.1.0.0.HpiLOAccount", 
            "LoginName": "support", 
            "Privileges": {
                "LoginPriv": True, 
                "RemoteConsolePriv": True, 
                "UserConfigPriv": True, 
                "VirtualMediaPriv": True, 
                "VirtualPowerAndResetPriv": True, 
                "iLOConfigPriv": True
            }, 
            "Type": "HpiLOAccount.1.0.0"
        }
    }, 
    "Password": None, 
    "Type": "ManagerAccount.1.0.0", 
    "UserName": "support", 
    "links": {
        "self": {
            "href": "/rest/v1/AccountService/Accounts/3"
        }
    }
}

def getResponse(requests):
    #{"UserName": "support", "Password": "11111111", "Oem": {"Hp": {"Privileges": {"RemoteConsolePriv": true, "iLOConfigPriv": true, "VirtualMediaPriv": true, "UserConfigPriv": true, "VirtualPowerAndResetPriv": true, "LoginPriv": true}, "LoginName": "support"}}}
    username = requests.get('UserName')
    password = requests.get("Password")
    res = createUserResp
    res['username'] = username
    #res['Password'] = password
    res['Oem']['Hp']['LoginName'] = username
    return res

