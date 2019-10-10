create table package_profile (
    id  varchar(64) primary key,
    pkg_name varchar(255) not null,
    pkg_desc varchar(255),
    repo_url varchar(255),
    pkg_version varchar(255) not null,
    author_name varchar(255),
    author_email varchar(255),
    author_desc varchar(255),
    pri_filename varchar(255),
    status varchar(2),
    created_at datetime,
    updated_at datetime
);
alter table package_profile add index idx_pkg_prf_name_ver(pkg_name, pkg_version);

create table last_version (
    id  varchar(64) primary key,
    pkg_name varchar(255) not null,
    pkg_version varchar(255) not null,
    pkg_profile_id varchar(64) not null,
    created_at datetime,
    updated_at datetime
);
alter table last_version add index idx_lst_ver_name(pkg_name);
alter table last_version add index idx_lst_ver_name_ver(pkg_name, pkg_version);
alter table last_version add index idx_lst_ver_pkg_prf_id(pkg_profile_id);
alter table last_version add index idx_lst_ver_crt_at(created_at);


create table upgrade_version (
    id  varchar(64) primary key,
    version varchar(255) not null,
    version_sort varchar(255) not null,
    description varchar(255),
    path varchar(255),
    status varchar(2),
    created_at datetime,
    updated_at datetime
);
alter table upgrade_version add index idx_upd_ver_ver(version);
alter table upgrade_version add index idx_upd_ver_ver_sort(version_sort);


create table console_user (
    id  varchar(64) primary key,
    username varchar(255) not null unique,
    password varchar(255) not null,
    status varchar(2),
    created_at datetime,
    updated_at datetime
);
