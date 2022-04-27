insert into public.users (id, username, password, created_at, updated_at)
values  (1, 'luqmansen', '1230j()U)GJIOPA', '2022-04-25 14:46:24.960559', '2022-04-25 14:46:24.960559'),
        (2, 'yeeterBoi', '1230j()U)GJIOPA', '2022-04-25 14:46:24.960559', '2022-04-25 14:46:24.960559'),
        (3, 'noggos', '1230j()U)GJIOPA', '2022-04-25 14:46:24.960559', '2022-04-25 14:46:24.960559');

SELECT setval('public."users_id_seq"',
              (SELECT MAX(id) FROM public.users)
           );

INSERT INTO public.posts (id, author_id, title, content, created_at, updated_at)
VALUES (1, 1, 'Test', 'Teest Content', '2022-04-25 14:49:16.633220', '2022-04-25 14:49:16.633220');

SELECT setval('public."posts_id_seq"',
              (SELECT MAX(id) FROM public.posts)
           );

insert into public.comments (id, parent_post_id, parent_id, author_id, content, created_at, updated_at)
values  (1, 1, null, 1, 'woii', '2022-04-25 14:49:51.379350', '2022-04-25 14:49:51.379350'),
        (2, 1, 1, 1, 'child 2', '2022-04-25 14:50:26.144688', '2022-04-25 14:50:26.144688'),
        (3, 1, 1, 1, 'child 1', '2022-04-25 14:50:33.190596', '2022-04-25 14:50:33.190596'),
        (4, 1, 3, 1, 'child heheh', '2022-04-25 14:50:40.381006', '2022-04-25 14:50:40.381006'),
        (5, 1, 4, 1, 'child again', '2022-04-25 14:50:49.071757', '2022-04-25 14:50:49.071757'),
        (6, 1, 5, 1, 'child nested again', '2022-04-25 14:50:56.528112', '2022-04-25 14:50:56.528112'),
        (7, 1, null, 1, 'child level 1 again', '2022-04-25 14:51:13.483050', '2022-04-25 14:51:13.483050');

SELECT setval('public."comments_id_seq"',
              (SELECT MAX(id) FROM public.comments)
           );

insert into public.reacts (id, name, thumbnail_url, created_at, updated_at)
values  (1, 'like', 'http://example.com/like.jpg', '2022-04-25 14:56:12.047897', '2022-04-25 14:56:12.047897'),
        (2, 'angery', 'http://example.com/angery.jpg', '2022-04-25 14:57:19.452581', '2022-04-25 14:57:19.452581');
