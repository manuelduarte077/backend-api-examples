package com.manuelduarte077.apirest.person;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class PersonService {

    private  final PersonRepository personRepository;

    public Person cratePerson(Person person){
        return personRepository.save(person);
    }

}
