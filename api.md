Things to do

- [x] Docker
- [x] First Setups and connection to DB

# Users
- [] Users table - email, password, age, firsname, lastname, gender
- [] routes to create, findall, findone, update and delete users
- [] route to login, authenticate with jwt and refreshtoken
- [] route to update recovery password

# Curriculum
- [] curriculum-user table - with about me - user_id -> user.id
- [] socials table - curriculum_id, link and title fields - curriculum_id -> curriculum.id
- [] languages - curriculum_id, language and level fields - curriculum_id -> curriculum.id
- [] experiences - curriculum_id, title, company, description, start_date, end_date - curriculum_id -> curriculum.id
- [] educations - curriculum_id, title, institution, description, start_date, end_date - curriculum_id -> curriculum.id
- [] skills - curriculum_id, title, description - curriculum_id -> curriculum.id
- [] projects - curriculum_id, title, description, link, start_date, end_date - curriculum_id -> curriculum.id
- [] routes to create, findall, findone, update and delete curriculum

# Curriculum IA
- [] route to generate about me using IA
- [] route to generate bullet points using IA

# Curriculum PDF
- [] route to generate pdf
- [] route to send pdf by email
