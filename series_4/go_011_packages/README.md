> Package Details

- Visibility
  - Private members available to all members regardless of file
  - Public/Exported members are accessible to imported file
  - Types of imports:
  
  Import Declaration | Local Name of Sin
  import "lib/math"    math.Sin
  import m "lib/math"  m.Sin
  import . "lib/math"  Sin
  
 
- Demo
  - Contrive Example Requirements
  1. Simple `awesome` application
  2. Simple  `security` package
        - Exposes: Login(), Logout(), IsAuthenticated()
        - Private: currentUser
  3. Recipe Manager Application
    - `recipe` package
    - Uses `secutiry` package