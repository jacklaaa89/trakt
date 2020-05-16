// package authorization
// The API uses OAuth2. If you know what's up with OAuth2, grab your library and starting rolling.
// If you have access to a web browser (mobile app, desktop app, website), use standard OAuth. If you don't
// have web browser access (media center plugins, smart watches, smart TVs, command line scripts, system services),
// use Device authentication.
//
// Application Flow:
//
// 1. Redirect to request Trakt access. Using the /oauth/authorize method, construct then redirect to this URL.
//    The Trakt website will request permissions for your app and the user will have the opportunity to sign up
//    for a new Trakt account or sign in with their existing account.
// 2. Trakt redirects back to your site. If the user accepts your request, Trakt redirects back to your site with
//    a temporary code in a code GET parameter as well as the state (if provided) in the previous step in a
//    state parameter. If the states don’t match, the request has been created by a third party and the process
//    should be aborted.
// 3. Exchange the code for an access token. If everything looks good in step 2, exchange the code for an access
//    token using the /oauth/token method. Save the access_token so your app can authenticate the user by
//    sending the Authorization header as indicated below or in any example code. The access_token is valid for
//    3 months. Save and use the refresh_token to get a new access_token without asking the user to re-authenticate.
package authorization
