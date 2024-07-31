package com.donmanuel.contactlistapi;


import com.donmanuel.contactlistapi.entity.Contact;
import com.donmanuel.contactlistapi.repository.ContactRepository;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.List;

@SpringBootApplication
public class ContactlistapiApplication {

    public static void main(String[] args) {
        SpringApplication.run(ContactlistapiApplication.class, args);

    }

    @Bean
    CommandLineRunner runner(
        ContactRepository contactRepository
    ) {
        return args -> {
            List<Contact> contacts = Arrays.asList(
                new Contact(
                    "John Doe", "manue@dev.co,", LocalDateTime.now()
                ),
                new Contact(
                    "John Doe", "manue@dev.co,", LocalDateTime.now()
                ),
                new Contact(
                    "John Doe", "manue@dev.co,", LocalDateTime.now()
                ),
                new Contact(
                    "John Doe", "manue@dev.co,", LocalDateTime.now()
                )
            );

            contactRepository.saveAll(contacts);
        };
    }


}
