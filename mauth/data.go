package mauth

var defaultResRolesText = `
# Auth Config

/Demo*=DEMO
/Admin*=ADMIN
/Root*=ROOT
/SysAdmin*=SysAdmin

# static assets
/public*=ALL
/mps/public*=ALL
/fav*=ALL
/robots*=ALL
/logo*=ALL
/assets*=ALL

# for old spring
/j_spring*=ALL

# some pages
/Help/*=ALL
/Links/*=ALL

/AppAjax/*=ALL
/Account*=ALL
/AccountAjax/=ALL
/All*=ALL
`
