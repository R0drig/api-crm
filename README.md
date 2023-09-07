/register POST
Register a user with 
{
    "email":"email@email.com",
    "name": "name"
    "passwd": "password"
}
/login POST
{
    "email":"email@email.com",
    "passwd: "Password"
}
Response JWT Token

/auth/user 
Get User info
{
    "Authorization": "KeyJwt"
}