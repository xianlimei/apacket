#!/usr/bin/env python
# -*- coding: UTF-8 -*-
body = '''<?xml version="1.0" encoding="utf-8" ?>
<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:ter="http://www.onvif.org/ver10/error" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:xs="http://www.w3.org/2000/10/XMLSchema" xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2">
<soap:Body>
<soap:Fault>
<soap:Code>
<soap:Value>soap:Sender</soap:Value><soap:Subcode>
<soap:Value>wsse:InvalidSecurity</soap:Value>
</soap:Subcode>
</soap:Code>
<soap:Reason>
<soap:Text xml:lang="en">An error was discovered processing the wsse:Security header.</soap:Text>
</soap:Reason>
<soap:Node>http://www.w3.org/2003/05/soap-envelope/node/ultimateReceiver</soap:Node>
<soap:Role>http://www.w3.org/2003/05/soap-envelope/role/ultimateReceiver</soap:Role>
</soap:Fault>
</soap:Body>
</soap:Envelope>'''

