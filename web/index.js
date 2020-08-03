import app from './app'

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
            ${app}
        </td>
        <td width="10" align="center" valign="top"></td>
    </tr>
</table>
<div id="footer"></div>
`

show_menu()
