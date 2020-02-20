ALTER TABLE users_appointments
    ADD CONSTRAINT users_appointments_user_id_users_id_foreign
    FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE users_appointments
    ADD CONSTRAINT users_appointments_appointment_id_appointments_id_foreign
    FOREIGN KEY (appointment_id) REFERENCES appointments (id) ON UPDATE CASCADE ON DELETE CASCADE;
