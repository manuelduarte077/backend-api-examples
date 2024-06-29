package com.donmanuel.contactlistapi.controller;


import com.donmanuel.contactlistapi.entity.Contact;
import com.donmanuel.contactlistapi.repository.ContactRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;


@RequestMapping("/api/contacts")
@RestController
public class ContactController {
    @Autowired
    private ContactRepository contactRepository;


    @GetMapping
    public Iterable<Contact> getContacts() {
        return contactRepository.findAll();
    }

    @GetMapping("/{id}")
    public Contact getContact(@PathVariable Integer id) {
        return contactRepository.findById(id).orElse(null);
    }

    @ResponseStatus(HttpStatus.CREATED)
    @PostMapping
    public Contact addContact(@RequestBody Contact contact) {
        contact.setCreatedAt(LocalDateTime.now());
        return contactRepository.save(contact);
    }

    @PutMapping
    public Contact updateContact(@PathVariable Integer id, @RequestBody Contact contact) {
        Contact contactToUpdate = contactRepository.findById(id).orElse(null);

        contactToUpdate.setName(contact.getName());
        contactToUpdate.setEmail(contact.getEmail());

        return contactRepository.save(contactToUpdate);
    }

    @ResponseStatus(HttpStatus.NO_CONTENT)
    @DeleteMapping("/{id}")
    public void deleteContact(Integer id) {

        Contact contactToDelete = contactRepository.findById(id).orElse(null);

        contactRepository.delete(contactToDelete);
    }
}
