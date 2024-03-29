// @ts-ignore
import { css, keyframes } from 'emotion'

const low = str => (str.charAt(0).toLowerCase() + str.substring(1))
const up = str => (str.charAt(0).toUpperCase() + str.substring(1))

export const s: {
  [key: string]: string,
} = {
  ...Object.fromEntries(
    ['p', 'm'].map(
      type => [-4,-2,-1,1,2,4,8,12,16].map(size => {
        const t = { p : 'padding', m: 'margin' }[type]
        return Object.entries({
          [`${type}${size}`]: css`${t}: ${size}px!important;`,
          [`${type}x${size}`]: css`${t}-left: ${size}px!important; ${t}-right: ${size}px!important;`,
          [`${type}y${size}`]: css`${t}-top: ${size}px!important; ${t}-bottom: ${size}px!important;`,
          [`${type}t${size}`]: css`${t}-top: ${size}px!important;`,
          [`${type}b${size}`]: css`${t}-bottom: ${size}px!important;`,
          [`${type}l${size}`]: css`${t}-left: ${size}px!important;`,
          [`${type}r${size}`]: css`${t}-right: ${size}px!important;`,
        })
      })
    ).flat(2)
  ),
  ...Object.fromEntries(
    ['h', 'v'].map(
      type => [-1,1,2,4,8,12,16].map(gap => {
        const margin  = { h: 'left', v: 'top' }[type]
        return [
          `${type}Gap${gap}`,
          css`
            & > * {margin-${margin}: ${gap}px!important;}
            & > *:first-child {margin-${margin}: 0!important;}
          `]
      })
    ).flat()
  ),
  ...Object.fromEntries(
    [2,4,8,12,16].map(size => [`bdRadius${size}`, css`border-radius: ${size}px!important;`])
  ),
  ...Object.fromEntries(
    ['', 'bd', 'bg'].map(type => {
      const t = { '': '', bd: 'border-', bg: 'background-' }[type]
      return Object.entries({
        [`${type}Primary`]: css`${t}color: #007bff!important;`,
        [`${type}Secondary`]: css`${t}color: #6c757d!important;`,
        [`${type}Success`]: css`${t}color: #28a745!important;`,
        [`${type}Danger`]: css`${t}color: #dc3545!important;`,
        [`${type}Warning`]: css`${t}color: #ffc107!important;`,
        [`${type}Info`]: css`${t}color: #17a2b8!important;`,
        [`${type}Light`]: css`${t}color: #f8f9fa!important;`,
        [`${type}Dark`]: css`${t}color: #343a40!important;`,
      }).map(([k, v]) => [low(k), v])
    }).flat()
  ),
  ...Object.fromEntries(
    [10,11,12,13,14,15,16,18,20].map(
      size => [`font${size}`, css`font-size: ${size}px!important;`]
    )
  ),
  ...Object.fromEntries(
    ['lighter','normal','bold','bolder'].map(
      weight => [`font${up(weight)}`, css`font-weight: ${weight}!important;`],
    )
  ),
  ...{
    dFlex: css`display: flex!important;`,
    flexColumn: css`flex-direction: column!important;`,
    flexRow: css`flex-direction: row!important;`,
    grow: css`flex-grow: 1!important;`,
    shrink: css`flex-shrink: 1!important;`,
    justifyStart: css`justify-content: start!important;`,
    justifyCenter: css`justify-content: center!important;`,
    justifyEnd: css`justify-content: end!important;`,
    justifyBetween: css`justify-content: space-between!important;`,
    alignStart: css`align-items: start!important;`,
    alignCenter: css`align-items: center!important;`,
    alignEnd: css`align-items: end!important;`,
    w100: css`width: 100%!important;`,
    // border
    bd: css`border: 1px solid #343a40!important;`,
    bdTop: css`border-top: 1px solid #343a40!important;`,
    bdBottom: css`border-bottom: 1px solid #343a40!important;`,
    bdLeft: css`border-left: 1px solid #343a40!important;`,
    bdRight: css`border-right: 1px solid #343a40!important;`,
  },
}

export const cardHeader = css`
background: linear-gradient(to bottom, #92a0a5  0%, #66757C 100%);
`

export const notificationContainer = css`
position: fixed;
width: 100%;
display: flex;
flex-direction: column;
top: 40px;
z-index: 1000;
`
export const notificationInstance = css`
margin: 16px auto auto;
&:first-child { margin-top: 0 }
padding: 8px 12px;
font-size: 12px;
box-shadow: 0 3px 6px -4px #333, 0 6px 16px 0 #333, 0 9px 28px 8px #333;
`

const loadingAnimation1 = keyframes`
0% {transform: scale(0);}
100% {transform: scale(1);}
`
const loadingAnimation2 = keyframes`
0% {transform: translate(0, 0);}
100% {transform: translate(20px, 0);}
`
const loadingAnimation3 = keyframes`
0% {transform: scale(1);}
100% {transform: scale(0);}
`
export const loadingContainer = css`
display: inline-block;
position: relative;
width: 52px;
height: 12px;
& div {
  position: absolute;
  top: 0;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #f8f9fa;
  animation-timing-function: cubic-bezier(0, 1, 1, 0);
}
& div:nth-child(1) {
  left: 0;
  animation: ${loadingAnimation1} 0.6s infinite;
}
& div:nth-child(2) {
  left: 0;
  animation: ${loadingAnimation2} 0.6s infinite;
}
}
& div:nth-child(3) {
  left: 20px;
  animation: ${loadingAnimation2} 0.6s infinite;
}
& div:nth-child(4) {
  right: 0;
  animation: ${loadingAnimation3} 0.6s infinite;
}
`

console.log(s)
