/*
Package platform implements internal third party services

DESIGN RULE: Package platform formats third party services with the
application's specifics like the following
    - email formats and templates
    - webhook dispatcher
    - SSE payload formats
    - worker

Platform provides third party services for /domain

Platform SHOULD NOT
    - be required for /domain to work
    - affect the flow of application
*/

package platform
