<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="X-UA-Compatible" content="IE=Edge">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta HTTP-EQUIV="Pragma" CONTENT="no-cache">
<meta HTTP-EQUIV="Expires" CONTENT="-1">
<link rel="shortcut icon" href="images/favicon.png">
<link rel="icon" href="images/favicon.png">
<title>Shadow</title>
<link rel="stylesheet" type="text/css" href="index_style.css">
<link rel="stylesheet" type="text/css" href="form_style.css">
<script src="https://cdn.jsdelivr.net/npm/react@16.13.1/umd/react.development.js"></script>
<script src="https://cdn.jsdelivr.net/npm/react-dom@16.13.1/umd/react-dom.development.js"></script>
<!--<script src="https://cdn.jsdelivr.net/npm/react@17.0.1/umd/react.production.min.js"></script>-->
<!--<script src="https://cdn.jsdelivr.net/npm/react-dom@16.13.1/umd/react-dom.production.min.js"></script>-->
<script type="text/javascript" src="/state.js"></script>
<script type="text/javascript" src="/general.js"></script>
<script type="text/javascript" src="/popup.js"></script>
<script type="text/javascript" src="/help.js"></script>
<script type="text/javascript" src="/validator.js"></script>
</head>
<body class="bg">

<div id="TopBanner"></div>
<div id="Loading" class="popup_bg"></div>
<table class="content" align="center" cellpadding="0" cellspacing="0">
<tr>
    <td width="17">&nbsp;</td>
    <td valign="top" width="202">
        <div id="mainMenu"></div>
        <div id="subMenu"></div>
    </td>
    <td valign="top">
        <div id="tabMenu" class="submenuBlock"></div>
        <table width="98%" border="0" align="left" cellpadding="0" cellspacing="0">
        <tr>
            <td valign="top">
            <table width="760px" border="0" cellpadding="4" cellspacing="0" bordercolor="#6b8fa3" class="FormTitle" id="FormTitle">
            <tr bgcolor="#4D595D">
                <td valign="top">
                    <div>&nbsp;</div>
                    <div class="formfonttitle">Tools - Shadow</div>
                    <div style="margin:10px 0 10px 5px;" class="splitLine"></div>
                    <div id="shadow-app"></div>
                </td>
            </tr>
            </table>
            </td>
        </tr>
        </table>
    </td>
    <td width="10" align="center" valign="top">&nbsp;</td>
</tr>
</table>
<div id="footer"></div>

<script>
  show_menu()

  function require(module) {
    return {
      'react': React,
      'react-dom': ReactDOM,
    }[module]
  }
</script>
<script type="text/javascript" src="/user/shadowApp.js" async></script>
</body>
</html>
