/*
  DESIGN RULE
  Package platform formats third party services with the application's
  specifics like the following
    [1] email formats and templates
    [2] webhook dispatcher
    [3] SSE payload formats
    [4] worker

  Platform provides third party services for domain package

  Platform SHOULD NOT
    [1] be required for domain to work
    [2] affect the flow of application
*/

package platform
