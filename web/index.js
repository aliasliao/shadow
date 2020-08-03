import app from "./app"

document.body.classList.toggle('bg')

const container = document.querySelector('#shadow-app')
container.innerHTML = `
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
                    <div class="formfonttitle">Shadow - airplane</div>
                    <div style="margin:10px 0 10px 5px;" class="splitLine"></div>
                    ${app}
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
`

show_menu()
