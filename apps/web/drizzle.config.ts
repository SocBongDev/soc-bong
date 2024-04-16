import type { Config } from 'drizzle-kit'

export default {
    schema: [
        './src/lib/permission/entities.ts',
        './src/lib/role/entities.ts',
        './src/lib/user/entities.ts',
        './src/lib/registration/entities.ts',
        './src/lib/role_permission/entities.ts',
        './src/lib/user_role/entities.ts',
        './src/lib/student/entities.ts',
        './src/lib/agency/entities.ts',
    ],
    out: './drizzle/migrations',
    driver: 'pg',
    dbCredentials: {
        host: process.env.DB_HOST,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME
    }
} satisfies Config
