ALTER TABLE users_appointments
    DROP CONSTRAINT users_appointments_user_id_users_id_foreign;

ALTER TABLE users_appointments
    DROP CONSTRAINT users_appointments_appointment_id_appointments_id_foreign;
