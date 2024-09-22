# Medchain-backend
Backend of Ethereum-based DApp for secure and transparent storage of medical prescriptions 

`````mermaid
---
title: DataBase Usuarios e Logs (1.0v)
---
classDiagram

   
    RolesUser  <|-- User
    AuditLogs <|-- User
    AccessLogs <|-- User
    RolesUser <|-- Roles
    RolesPermissions <|-- Roles
    RolesPermissions <|-- Permissions

   
    class User{
        int idUser_pk PK 
        String name NOT NULL
        String email NOT NULL
        String password NOT NULL
    }
   
    class Permissions{
        int idPermissions_pk PK
        String name NOT NULL
        String description

   
    }
    class Roles{
        int idRoles_pk PK
        String name NOT NULL
        String description
    }
     class RolesUser{
        int idUser_pk_fk FK PK
        int idRoles_pk_fk FK PK
    }
    class RolesPermissions{
        int idRoles_pk_fk FK PK
        int idPermissions_pk_fk FK PK
    }
    class AuditLogs {
        int idAuditLogs_pk PK,
        String table_name 
        String operation_type -- INSERT, UPDATE e DELETE 
        int id_record
        int idUser_fk FK  
        JSONB old_values 
        JSONB new_values 
        DateTime date 
    
    }
   


`````
