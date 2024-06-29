package com.donmanuel.contactlistapi.repository;

import com.donmanuel.contactlistapi.entity.Contact;
import org.springframework.data.repository.CrudRepository;

public interface ContactRepository extends CrudRepository<Contact, Integer> {



}
