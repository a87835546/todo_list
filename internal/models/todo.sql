create table task
(
    id            bigint auto_increment
        primary key,
    user_id       bigint                      not null,
    type          int           default 0  not null comment 'task type  0',
    task_group_id int           default 0  not null,
    title         varchar(1000)            not null,
    `desc`        varchar(2000) default '' not null,
    process       int           default 0  not null comment '0-100',
    end_time      bigint        default 0  not null,
    start_time    bigint        default 0  not null,
    project_name  varchar(200)  default '' not null,
    created_at    bigint        default 0  not null,
    updated_at    bigint        default 0  null
);

create table task_group
(
    id         bigint auto_increment
        primary key,
    color      bigint       default 4294967295 not null comment '0xffffffff',
    icon       varchar(200) default ''         not null,
    name       varchar(100) default ''         not null,
    name_en    varchar(200) default ''         not null,
    user_id    bigint          default 0          not null comment 'if who created new group then this colum will not be 0',
    created_at bigint       default 0          not null,
    updated_at bigint       default 0          not null
);

create table user
(
    id           bigint auto_increment
        primary key,
    username     varchar(200)  default '' not null,
    email        varchar(200)  default '' not null,
    password     varchar(1000) default '' not null,
    is_delete    int           default 0  not null,
    account_type int           default 0  not null,
    created_at   bigint        default 0  not null,
    updated_at   bigint        default 0  not null,
    avatar       varchar(500)  default '' not null,
    login_ip     varchar(100)  default '' not null,
    register_ip  varchar(100)  default '' not null,
    device_type  int           default 0  not null,
    constraint id_index
        unique (id)
);

