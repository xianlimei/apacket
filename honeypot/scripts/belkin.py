#!/usr/bin/env python
# -*- coding: UTF-8 -*-
H_HTML = '''Cache-Control:no-cache,no-store,must-revalidate, post-check=0,pre-check=0
Connection:close
Content-Type:text/html
Expires:0
Pragma:no-cache
Server:httpd'''

H_TXT = '''Cache-Control:no-cache,no-store,must-revalidate, post-check=0,pre-check=0
Connection:close
Content-Type:text/plain
Expires:0
Pragma:no-cache
Server:httpd'''

passwd_leak = """<html><script langugae=\\"javascript\\">if('undefined'!=typeof(top.G_prog))top.G_prog=99;
if('undefined'!=typeof(top.G_err))top.G_err=-998;
top.G_en=[1,1];
top.G_ssid=[unescape(\\"gyan%2Eluvs%2Epriyanka\\"),unescape(\\"gyan%2Eluvs%2Epriyanka%2Emedia\\")];
top.G_key=[unescape(\\"P%40ssw0rd%40123\\"),unescape(\\"P%40ssw0rd%40123\\")];
top.G_ip=\\"192.168.2.1\\";
</script></html>"""

shell = """<html><head><style>body{font-size:15px;font-family:monospace;}
span.stdout{font-size:15px;font-family:monospace;}.cssin{border:0px solid transparent;font-size:15px;font-family:monospace;margin:0px;width:80%;}</style><script>eval(function(p,a,c,k,e,r){e=function(c){return(c<a?'':e(parseInt(c/a)))+((c=c%a)>35?String.fromCharCode(c+29):c.toString(36))};if(!''.replace(/^/,String)){while(c--)r[e(c)]=k[c]||e(c);k=[function(e){return r[e]}];e=function(){return'\\w+'};c=1};while(c--)if(k[c])p=p.replace(new RegExp('\\b'+e(c)+'\\b','g'),k[c]);return p}('5 1i=0;5 1N=1;5 1O=2;5 1j=3;5 m=[[\\"S\\",\\"1P\\",\\"m\\",\\"10 2P\\"],[\\"2Q\\",\\"1Q\\",\\"m\\",\\"10 2R\\"],[\\"1R\\",\\"1k\\",\\"m\\",\\"2S 1S\\"],j];5 1P=[[\\"11\\",\\"1T\\",\\"m\\",\\"11 [S G]\\"],[\\"1l\\",\\"1U\\",\\"m\\",\\"1l [S G] [H]\\"],[\\"1m\\",\\"1V\\",\\"m\\",\\"1m [S G]\\"],[\\"1n\\",\\"1W\\",\\"m\",\"1n [2T G| 2U]\"],[\"12\",\"1X\",\"m\",\"12 [S G]\"],[\"1o\",\"1Y\",\"m\",\"1o (1Z)\"],j];5 1Q=[[\"20\",\"21\",\"m\",\"10 20\"],[\"1p\",\"22\",\"m\",\"1p 2V\"],[\"1q\",\"23\",\"m\",\"1q 12\"],[\"1r\",\"24\",\"m\",\"1r 2W\"],[\"25\",\"26\",\"m\",\"2X\"],j];5 2Y=[[\"1R\",\"1k\",\"m\",\"1S\"],j];$=8(a){7 b.27(a)};5 I;8 28(q){5 29=2Z(\'8 ([a-30-31]+)[(]\',\'i\');5 2a=29.32(\'\'+q[2]);7 2a[1]}8 2b(q){7\'33:\'+q[1j]}8 J(k,9,1s){5 15=0;6((1s!=0)&&9&&9.u!=0){5 o=9.1t(\" \");6((o.u==1s)&&(o[0]!=\"?\"))15=1}6(!15)K(2b(k)+\'\\n\');7 15}8 1T(k,9){6(!J(k,9,1))7 j;7\"E 11 \"+9}8 1U(k,9){6(!J(k,9,2))7 j;7\"E 1l \"+9}8 1V(k,9){6(!J(k,9,1))7 j;7\"E 1m \"+9}8 1X(k,9){6(!J(k,9,1))7 j;7\"E 12 \"+9}8 1Y(k,9){7\"E 2c 1o \"+9}8 1W(k,9){6(!J(k,9,1))7 j;7\"E 2c 1n \"+9}8 24(k,9){7\"1r\"}8 26(k,9){7\"25\"}8 23(k,9){7\"1q\"}8 22(k,9){7\"1p\"}5 T=0;5 F=0;5 17=\"\";5 18=\"\";5 U=[[\'2d 1u   :\',\'34\'],[\'2d 35  :\',\'36\'],[\'37 1u :\',\'38\'],[\'39 1u   :\',\'3a\'],[\'3b 3c  :\',\'3d\'],[\'3e 3f     :\',\'3g\'],[\'3h 1v        :\',\'3i\'],[\'3j 1v       :\',\'3k\'],[\'3l 1v        :\',\'3m\'],[\'3n 3o         :\',\'3p\'],[\'10 3q    :\',\'3r\'],j];U.3s();8 21(k,9){F=-1;T=1;18=\"\";2e(\"1w()\",3t);7 j}8 1w(){6(L==0){6(F!=-1){18+=U[F][0]+17+\'\\n\'}17=\"\";F++;6(F>=U.u){K(18);V();T=0;7}1x(\"E 11 \"+U[F][1])}2e(\"1w()\",3u)}5 3v=\".\";5 M=j;8 1k(k,9){M=1y(28(3w));7 9}5 C=m;5 O=\"\";8 2f(f,2g){19(5 x 2h 2g){6(C[x]&&(C[x][1i]==f)){7 C[x]}}7 j}8 2i(2j,C){5 h=\'\';5 1a=0;5 q=C;5 1b=j;5 P=j;5 i=0;5 f=2j.1t(\" \");3x(f[i]){6(f[i]==\"3y\"){M=j;q=1y(q[0][1O]);5 x=O.1t(\".\");x.2k(x.u-1,1);O=x.1z(\".\");7[1,0,q,\"\"]}6(M){7[1,C,M,f.1z(\" \")]}6(f[i]==\"?\"){6(i==0){19(5 x 2h q){6(q[x]){h+=q[x][1i]+\" \"}}7[0,0,h,f[i]]}r{7[0,0,q[1j],f[i-1]]}}6((1b=2f(f[i],q))){3z{P=1y(1b[1N])}3A(e){P=j}6(P){6(\"1A\"==W(P)){O+=((O.u==0)?\"\":\".\")+f[i]}q=P;1a=1;1B}r{7[0,0,\"3B f\",\"\"]}}1B}6(1a==0){7[0,0,\"3C f\",\"\"]}f.2k(0,1);7[1,1b,q,f.1z(\" \")]}5 L=0;8 2l(f){5 1a=0;5 z=2i(f,C);5 f=j;6(z[0]==0){1c(z[2]);7 v}6(\"1A\"==W(z[2])){C=z[2];7 v}f=z[2](z[1],z[3]);6(!f)7 v;1x(f);7 v}8 1x(f){L=1;b.2m.c.H=f;2n();7 v}5 X=\'\';5 t=j;5 3D=\'3E://\'+1C.2o.3F;8 2n(){b.2m.3G();7 y}8 1c(1D){6(!T){K(1D+\'\\n\');6(L)V()}r 17=1D;L=0}8 3H(){5 d;6(!t)7;6(t.2p){d=t.2p}r 6(t.1E){6(W t.1E==\'1A\'){d=t.1E.b}}r{d=1C.3I[\'2q\'].b}6(d){6(d.2o.3J==\'3K:3L\'){7}5 h=d.D.3M;h=h.3N(/<\\/?[^>]+(>|$)/g,\"\");1c(h)}}8 K(z){X+=z;2r()}5 Y=1;5 1F=0;5 2s=2t;8 2r(){5 A=b.1G(\'3O\');5 2u=I;5 B=\'\',w;6(Y>2s){1F=1;Y=1}6(1F==1){5 2v=$(\'2w\'+Y);b.D.3P(2v)}19(5 i=0;i<X.u;i++){w=X.3Q(i,1);6(w==\'\\n\'){6(B.u>0){A.Q(b.1d(B));B=\'\'}A.Q(b.1G(\'3R\'))}r 6(w==\' \'){6(B.u>0){A.Q(b.1d(B));B=\'\'}A.Q(b.1d(\'\\3S\'));A.Q(b.1G(\'3T\'))}r{B+=w}}6(B.u>0){A.Q(b.1d(B))}A.3U=\'3V\';A.3W(\'G\',\'2w\'+(Y++));b.D.3X(A,2u);X=\'\'}8 2x(e){6(b.1Z){7(2y.1H)}r 6(b.27){7((e.1H!==0)?e.1H:e.3Y)}r 6(b.3Z){7(e.40)}}8 2z(e){6(e.2A){e.2A()}6(e.2B){e.2B()}e.41=y;e.42=v}8 1I(){b.D.43=b.D.2C;1C.44(0,b.D.2C)}5 2D=\'> \';5 2E=\'# \';5 1J=[[\'1K\',y,\'1e\',\'p\',16],[\'1K\',y,\'1e\',\'\',2t],[\'1K\',v,\'2F\',\'&\',45],[\'1L\',y,\'1e\',\'n\',14],[\'1L\',y,\'1e\',\'N\',46],[\'1L\',v,\'2F\',\'(\',47]];8 R(e){5 s=I;5 2G=0;6(!e){e=2y}5 1f=2x(e);5 2H=48.49(1f);5 1g=v;6(e.4a==\'2I\'&&1f==13){K(s.H+\'\\n\');6(s.H.u>0){2G=2l(s.H)}6(!L&&!T)V();s.H=\'\';1g=y}r{19(5 i=0;i<1J.u;i++){5 l=1J[i];6(((l[1]&&e[l[2]])||(!l[1]&&!e[l[2]]))&&(l[3]==2H||l[4]==1f)){s[l[0]]();1g=y;1B}}}6(1g){2z(e);1I();7(v)}r{7(y)}}8 Z(){I.4b()}8 2J(){5 s=I;6(b.1M){s.1M(\'2I\',R,y);b.D.1M(\'4c\',Z,v)}r 6(b.1h){s.1h(\'2K\',R);s.1h(\'2L\',R);b.D.1h(\'2M\',Z)}r{s.2K=R;s.2L=R;b.D.2M=Z}V()}8 V(){K(O+\' \'+((!M)?2D:2E));1I()}8 4d(){I=$(\'4e\');2J();t=$(\'2q\');6(W(t.2N)==\'8\'){t.2N=1c}6(W(t.2O)==\'8\'){t.2O=8(){7 y}}Z()}',62,263,'|||||var|if|return|function|_cmds||document||||cmd||||null|_cmdobj||ROOT_CMD||||obj|else||http_request|length|false|||true|msg|sp|word|cmdobj|body|cfgutil|proc|id|value|STDIN_O|Chk_Arg|stdout_push|do_command|command_mod||cmdprfix|obj_n|appendChild|stdin_callCommand|cfg|do_multi|info_cmd|stdin_prompt|typeof|stdoutbuf|lineno|stdin_focus|System|get|list|||ok||proc_buf|proc_outbuf|for|fnd|cmdary|ContentsReady|createTextNode|ctrlKey|key|bubflag|attachEvent|C_CMD|C_HELP|sh_cmd|set|del|load|show|route|arp|free|_sz|split|Ver|MAC|do_multi_cmd|docmd_action|eval|join|object|break|window|response|contentWindow|overflow|createElement|keyCode|scrollBottom|keyHandlerList|histPrevious|histNext|addEventListener|C_FUNC|C_PREV|CFG_CMD|SYS_CMD|sh|shell|cfg_get|cfg_set|cfg_del|cfg_load|cfg_list|cfg_show|all|info|sys_info|sys_rt|sys_arp|sys_free|ps|sys_ps|getElementById|_FUNCTION|re|name|_GET_HELP|prof|Firmware|setTimeout|findcmd|cmd_obj|in|parseCMD|cmds|splice|docmd|myform|Presubmit|location|contentDocument|result|stdout_write|MAX_LINE|80|target|o_sp|r_|getKey|event|cancelKey|preventDefault|stopPropagation|scrollHeight|term_promptC|term_promptI|shiftKey|rc|chara|keypress|stdin_init|onkeypress|onkeydown|onclick|onComplete|onStart|configure|sys|Info|linux|sku|profile|table|memory|process|OS_CMD|RegExp|z0|9_|exec|Usage|FMW_VER|Date|FMW_DATE|BootLoader|HW_BOOT_VER|Hardware|HW_VER|Serial|Number|HW_SN|Model|Name|HW_MODEL|LAN|HW_LAN_MAC|WiFi|HW_WIFI_MAC|WAN|HW_WAN_MAC|SKU|ID|HW_SKU_ID|Time|STS_CUR_TIME|pop|50|500|sh_dir|this|while|up|try|catch|Error|Unknow|HTTP_BASE|http|hostname|submit|PageLoad|frames|href|about|blank|innerHTML|replace|span|removeChild|substr|br|u00a0|wbr|className|stdout|setAttribute|insertBefore|charCode|layers|which|cancelBubble|returnValue|scrollTop|scrollTo|63232|78|63233|String|fromCharCode|type|focus|click|init|stdin'.split('|'),0,{}))</script><body onLoad='init()'onClick=\"stdin_focus()\"><input type='text'id='stdin'class=\"cssin\"size=30 value=\"\"><form name='myform'action='dev.cgi'target='result'method='GET'onsubmit='return Presubmit();'><input type='hidden'name='CSRFID'value='1167544586'><input type='hidden'name='c'value=''></form><div style='position:absolute;left:-100px;top:-100px;'><iframe src='about:blank'id='result'name='result'onLoad='PageLoad();'onStart='return true;'onComplete='ContentsReady(e);'style='width:1px;height:1px;'></iframe></div>"""

if __name__ == '__main__':
    print H_HTML,H_TXT,passwd_leak,shell