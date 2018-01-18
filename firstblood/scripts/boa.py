#!/usr/bin/env python
# -*- coding: UTF-8 -*-
boa = '''<!-- <!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd"> -->
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=EUC-KR">
<title>웹 로그인</title>

<link href="/css/redmond/jquery-ui-1.9.2.custom.min.css" rel="stylesheet" /><link href="/css/redmond/custom.css" rel="stylesheet" /><!--[if IE 8]><link href="/css/redmond/custom_ie8.css" rel="stylesheet" /><![endif]--><script src="/js/jquery-1.8.3.js"></script><script src="/js/jquery-ui-1.9.2.custom.min.js"></script><script src="/js/common.js"></script>
<script src="../js/login.js"></script>
<script language="JavaScript">
<!--
	/*
	 *	HTML and JavaScript developed by Jake Jun.
	 *  email: hyun924@gmail.com 
	 *  facebook: www.facebook.com/hyun924
	 */
		
-->	 
</script>

<style>

</style>

</head>

<body leftmargin="0" topmargin="0" rightmargin="0" bottommargin="0" marginwidth="0" marginheight="0" onload="onload_body();" onunload="onunload_body();">
<form name="login_wconf" method="post" action="./login_proc.cgi">
	<input type=hidden id="login_os" name="login_os" value="win">
	<input type=hidden id="login_type" name="login_type" value="1">
	
<!--	
	<table width="100%" border="0" cellspacing="0" cellpadding="0" bgcolor="#6FA7D1">
		<tr>
			<td align="right">
				<table border="0" cellspacing="0" cellpadding="3px">
                    <tr>
                        <td valign="center" align="right" class="user">
                            <img src="../img/language.gif" height="12px"/>&nbsp;##RES_LANGUAGE##&nbsp; 
                        </td>
                        <td valign="center" align="center" width="150px">
							<select id="select_language" class="select_lang">
                                ##PARAM_LANGUAGE##
                            </select> 
                        </td>
                    </tr>
                </table>
			</td>
		</tr>
	</table>
-->	
	<table width="100%" height="100%" background="../img/login_bg.gif" border="0" cellspacing="0" cellpadding="0">
		<tr><td align="center" valign="center">	
			<table border="0" cellspacing="0" cellpadding="0">
				<tr>
					<td width="10px">&nbsp;</td>
					<td>
			  			<div id="tabs" class="tabs">
			                <ul>
			                    <li>
			                        <a href="#tabs-1"><div class="tab_header">웹 로그인</div></a>
			                    </li>
			                </ul>
			                <div id="tabs-1">
			                	<center>
			                		<table border="0" cellspacing="0" cellpadding="0px">
				                    	<tr><td height="3px">&nbsp;</td></tr>
									</table>
									
									<table border="0" cellspacing="0" cellpadding="0px">
										<tr>
											<td align="right" valign="center">
												<img src="../img/login_img.gif" />
											</td>
											<td align="center" valign="center" width="20px">
												&nbsp;
											</td>
											<td align="center" valign="center">
												<table border="0" cellspacing="0" cellpadding="0px">
							                    	<tr>
														<td align="right" valign="center" class="user">
															사용자 아이디&nbsp;
														</td>
														<td valign="center" class="user">
															&nbsp<input id="login_id" name="login_id" class="i_text01" size="15" type="text" onKeyPress="if(event.keyCode == \'13\') check_form();">
														</td>
														<td rowspan="2" valign="center" align="center">
															
														</td>
							                    	</tr>
													<tr>
							                    		<td align="right" valign="center" class="user">
							                    			비밀번호&nbsp;
														</td>
														<td valign="center" class="user">
															&nbsp<input id="login_pwd" name="login_pwd" class="i_text01" size="15" type="password" width="100px" onKeyPress="if(event.keyCode == \'13\') check_form();">
														</td>
													</tr>
							                    </table>
											</td>
											<td align="center" valign="center" width="20px">
												&nbsp;
											</td>
											<td align="left" valign="center">
												<a id="submit" onClick="check_form();" tabindex="-1">로그인</a>
											</td>
				                    	</tr>
										<tr height="20px">
											<td colspan="5"></td>
										</tr>
										<tr height="34px">
											<td colspan="5" valign="center" align="center" bgcolor="#A6C9E2">
												<div class="user">
													&nbsp;&nbsp;&nbsp;보안 관계상 승인된 사용자만 웹 서버에 접속을 할 수 있습니다.&nbsp;&nbsp;&nbsp;
												</div>
											</td>
										</tr>
									</table>
				                    
								</center>
			                </div>
			            </div>
					</td>
					<td width="10px">&nbsp;</td>
				</tr>
			</table>
			
		</td></tr>
	</table>
</form>
</body>

<script language="JavaScript">
<!--

-->
</script>

</html>'''
