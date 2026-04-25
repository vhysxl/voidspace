task
1. Create UI with dummy and make sure it looks good
2. Connect to backend and make sure it works
3. Add more features (Iterative process)


4/6/2026
1. created few pages and components
2. Completed Explore page with interactive tabs (Posts, Comments, Users).
3. Implemented Light/Dark mode theme with Zustand persistence and CSS variable mapping.
4. Created Settings page with theme toggle and secure account deletion workflow.
5. Implemented "New Post" modal and integrated it with desktop/mobile navigation.
6. Refactored API calls to BFF (Backend-for-Frontend) pattern to hide x-api-key from the browser.
7. Fixed critical bug in User Microservice where registration returned User ID "0", causing subsequent 404s.
8. Standardized layout by making Explore and Settings pages full-width.

Current Status:
- Removed theme switching logic; implemented permanent dark mode for the entire application.
- Standardized layout by making Explore and Settings pages full-width.

Task for tomorrow:
1. Ensure all components correctly use the background/foreground variables in the fixed dark mode.
2. Continue connecting remaining UI components to the backend.
