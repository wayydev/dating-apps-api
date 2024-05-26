PGDMP         .                |         
   datingapps %   14.11 (Ubuntu 14.11-0ubuntu0.22.04.1)    15.3 ?    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16548 
   datingapps    DATABASE     r   CREATE DATABASE datingapps WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C.UTF-8';
    DROP DATABASE datingapps;
                postgres    false            �           0    0    DATABASE datingapps    ACL     0   GRANT ALL ON DATABASE datingapps TO datingapps;
                   postgres    false    4275                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                postgres    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   postgres    false    5            �           0    0    SCHEMA public    ACL     Q   REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;
                   postgres    false    5            �            1259    19194    notifications    TABLE     7  CREATE TABLE public.notifications (
    id bigint NOT NULL,
    user_id bigint,
    title character varying(100),
    message character varying(356),
    data json,
    is_read boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
 !   DROP TABLE public.notifications;
       public         heap 
   datingapps    false    5            �            1259    19193    notifications_id_seq    SEQUENCE     }   CREATE SEQUENCE public.notifications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.notifications_id_seq;
       public       
   datingapps    false    5    226            �           0    0    notifications_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;
          public       
   datingapps    false    225            �            1259    19147    package_features    TABLE       CREATE TABLE public.package_features (
    id bigint NOT NULL,
    package_id bigint,
    name character varying(50),
    value text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
 $   DROP TABLE public.package_features;
       public         heap 
   datingapps    false    5            �            1259    19146    package_features_id_seq    SEQUENCE     �   CREATE SEQUENCE public.package_features_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.package_features_id_seq;
       public       
   datingapps    false    5    218            �           0    0    package_features_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.package_features_id_seq OWNED BY public.package_features.id;
          public       
   datingapps    false    217            �            1259    19138    packages    TABLE     �   CREATE TABLE public.packages (
    id bigint NOT NULL,
    name character varying(50),
    price bigint,
    description text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.packages;
       public         heap 
   datingapps    false    5            �            1259    19137    packages_id_seq    SEQUENCE     x   CREATE SEQUENCE public.packages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.packages_id_seq;
       public       
   datingapps    false    216    5            �           0    0    packages_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.packages_id_seq OWNED BY public.packages.id;
          public       
   datingapps    false    215            �            1259    19203    swaps    TABLE     �   CREATE TABLE public.swaps (
    id bigint NOT NULL,
    user_id bigint,
    swap_user_id bigint,
    "like" boolean,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.swaps;
       public         heap 
   datingapps    false    5            �            1259    19202    swaps_id_seq    SEQUENCE     u   CREATE SEQUENCE public.swaps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.swaps_id_seq;
       public       
   datingapps    false    5    228            �           0    0    swaps_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.swaps_id_seq OWNED BY public.swaps.id;
          public       
   datingapps    false    227            �            1259    19173    user_packages    TABLE       CREATE TABLE public.user_packages (
    id bigint NOT NULL,
    user_id bigint,
    package_id bigint,
    expires_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
 !   DROP TABLE public.user_packages;
       public         heap 
   datingapps    false    5            �            1259    19172    user_packages_id_seq    SEQUENCE     }   CREATE SEQUENCE public.user_packages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.user_packages_id_seq;
       public       
   datingapps    false    222    5            �           0    0    user_packages_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.user_packages_id_seq OWNED BY public.user_packages.id;
          public       
   datingapps    false    221            �            1259    19180    user_photos    TABLE     �   CREATE TABLE public.user_photos (
    id bigint NOT NULL,
    user_id bigint,
    photo text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.user_photos;
       public         heap 
   datingapps    false    5            �            1259    19179    user_photos_id_seq    SEQUENCE     {   CREATE SEQUENCE public.user_photos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.user_photos_id_seq;
       public       
   datingapps    false    224    5            �           0    0    user_photos_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.user_photos_id_seq OWNED BY public.user_photos.id;
          public       
   datingapps    false    223            �            1259    19161    users    TABLE       CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    username text,
    email text NOT NULL,
    password text,
    phone character varying(20),
    latitude numeric,
    longitude numeric,
    location public.geography(Point,4326),
    place_of_birth character varying(50),
    date_of_birth date,
    gender text,
    photo_profile text,
    find_on_distance bigint DEFAULT 20,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);
    DROP TABLE public.users;
       public         heap 
   datingapps    false    5    5    5    5    5    5    5    5    5    5    5    5    5    5    5    5    5            �            1259    19160    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public       
   datingapps    false    220    5            �           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public       
   datingapps    false    219            �           2604    19197    notifications id    DEFAULT     t   ALTER TABLE ONLY public.notifications ALTER COLUMN id SET DEFAULT nextval('public.notifications_id_seq'::regclass);
 ?   ALTER TABLE public.notifications ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    226    225    226            �           2604    19150    package_features id    DEFAULT     z   ALTER TABLE ONLY public.package_features ALTER COLUMN id SET DEFAULT nextval('public.package_features_id_seq'::regclass);
 B   ALTER TABLE public.package_features ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    218    217    218            �           2604    19141    packages id    DEFAULT     j   ALTER TABLE ONLY public.packages ALTER COLUMN id SET DEFAULT nextval('public.packages_id_seq'::regclass);
 :   ALTER TABLE public.packages ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    215    216    216            �           2604    19206    swaps id    DEFAULT     d   ALTER TABLE ONLY public.swaps ALTER COLUMN id SET DEFAULT nextval('public.swaps_id_seq'::regclass);
 7   ALTER TABLE public.swaps ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    227    228    228            �           2604    19176    user_packages id    DEFAULT     t   ALTER TABLE ONLY public.user_packages ALTER COLUMN id SET DEFAULT nextval('public.user_packages_id_seq'::regclass);
 ?   ALTER TABLE public.user_packages ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    221    222    222            �           2604    19183    user_photos id    DEFAULT     p   ALTER TABLE ONLY public.user_photos ALTER COLUMN id SET DEFAULT nextval('public.user_photos_id_seq'::regclass);
 =   ALTER TABLE public.user_photos ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    224    223    224            �           2604    19164    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public       
   datingapps    false    220    219    220            �          0    19194    notifications 
   TABLE DATA           w   COPY public.notifications (id, user_id, title, message, data, is_read, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    226   zI       �          0    19147    package_features 
   TABLE DATA           k   COPY public.package_features (id, package_id, name, value, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    218   �I       �          0    19138    packages 
   TABLE DATA           d   COPY public.packages (id, name, price, description, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    216   �I       �          0    17995    spatial_ref_sys 
   TABLE DATA           X   COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
    public          postgres    false    211   oJ       �          0    19203    swaps 
   TABLE DATA           f   COPY public.swaps (id, user_id, swap_user_id, "like", created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    228   �J       �          0    19173    user_packages 
   TABLE DATA           p   COPY public.user_packages (id, user_id, package_id, expires_at, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    222   �J       �          0    19180    user_photos 
   TABLE DATA           ]   COPY public.user_photos (id, user_id, photo, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    224   �K       �          0    19161    users 
   TABLE DATA           �   COPY public.users (id, name, username, email, password, phone, latitude, longitude, location, place_of_birth, date_of_birth, gender, photo_profile, find_on_distance, created_at, updated_at, deleted_at) FROM stdin;
    public       
   datingapps    false    220   �K       �           0    0    notifications_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.notifications_id_seq', 1, false);
          public       
   datingapps    false    225            �           0    0    package_features_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.package_features_id_seq', 2, true);
          public       
   datingapps    false    217            �           0    0    packages_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.packages_id_seq', 2, true);
          public       
   datingapps    false    215            �           0    0    swaps_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.swaps_id_seq', 1, false);
          public       
   datingapps    false    227            �           0    0    user_packages_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('public.user_packages_id_seq', 14, true);
          public       
   datingapps    false    221            �           0    0    user_photos_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.user_photos_id_seq', 1, false);
          public       
   datingapps    false    223            �           0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 14, true);
          public       
   datingapps    false    219            	           2606    19201     notifications notifications_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.notifications DROP CONSTRAINT notifications_pkey;
       public         
   datingapps    false    226            �           2606    19154 &   package_features package_features_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.package_features
    ADD CONSTRAINT package_features_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.package_features DROP CONSTRAINT package_features_pkey;
       public         
   datingapps    false    218            �           2606    19145    packages packages_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.packages
    ADD CONSTRAINT packages_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.packages DROP CONSTRAINT packages_pkey;
       public         
   datingapps    false    216                       2606    19208    swaps swaps_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.swaps
    ADD CONSTRAINT swaps_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.swaps DROP CONSTRAINT swaps_pkey;
       public         
   datingapps    false    228                       2606    19178     user_packages user_packages_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.user_packages
    ADD CONSTRAINT user_packages_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.user_packages DROP CONSTRAINT user_packages_pkey;
       public         
   datingapps    false    222                       2606    19187    user_photos user_photos_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.user_photos
    ADD CONSTRAINT user_photos_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.user_photos DROP CONSTRAINT user_photos_pkey;
       public         
   datingapps    false    224                       2606    19171    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public         
   datingapps    false    220                       2606    19169    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public         
   datingapps    false    220                       2606    19155 ,   package_features fk_packages_package_feature    FK CONSTRAINT     �   ALTER TABLE ONLY public.package_features
    ADD CONSTRAINT fk_packages_package_feature FOREIGN KEY (package_id) REFERENCES public.packages(id);
 V   ALTER TABLE ONLY public.package_features DROP CONSTRAINT fk_packages_package_feature;
       public       
   datingapps    false    216    218    4093                       2606    19214    swaps fk_swaps_swap_user    FK CONSTRAINT     |   ALTER TABLE ONLY public.swaps
    ADD CONSTRAINT fk_swaps_swap_user FOREIGN KEY (swap_user_id) REFERENCES public.users(id);
 B   ALTER TABLE ONLY public.swaps DROP CONSTRAINT fk_swaps_swap_user;
       public       
   datingapps    false    220    228    4099                       2606    19209    swaps fk_swaps_user    FK CONSTRAINT     r   ALTER TABLE ONLY public.swaps
    ADD CONSTRAINT fk_swaps_user FOREIGN KEY (user_id) REFERENCES public.users(id);
 =   ALTER TABLE ONLY public.swaps DROP CONSTRAINT fk_swaps_user;
       public       
   datingapps    false    228    220    4099                       2606    19188    user_photos fk_users_user_photo    FK CONSTRAINT     ~   ALTER TABLE ONLY public.user_photos
    ADD CONSTRAINT fk_users_user_photo FOREIGN KEY (user_id) REFERENCES public.users(id);
 I   ALTER TABLE ONLY public.user_photos DROP CONSTRAINT fk_users_user_photo;
       public       
   datingapps    false    4099    224    220            �      x������ � �      �   P   x�3�4�����,�44�4202�50�52S02�21�24�374143�6 $��e�i5K��jKCSc�f�e�f��qqq ���      �   h   x�3�t+JM�4�����,QpLNN-.�4202�50�52S02�2��25�32�01��60  ��e�P���Y��ih ��y9 �SS�oafjf��|�,��=... �:)�      �      x������ � �      �      x������ � �      �   �   x���ˍ�0E�5]���#�_����1�.�Pػk��S���!�-�F��b\���K�z.2�ʄ�C����^�Q҃��޻�Σ�cdǔ}�[�|�P���x|��HO�G�RML)3g��4j'�s8:�5�N�࠯�裦Σ�3h�<T���y�p�кS�Wm�s�Q�sN�zl�L=[�V���7��(K�@r6����t�j]���;muI1�|�:5N�Xm���jH7_������      �      x������ � �      �   z  x���Is�<�׾��w�G�U�y���6�3�׷d2bW���AJ�T��{^Ak��W��1�g����=۾Z�����E<�v��&v��C�� �w��sM�0����r��z�>��}����C�
!0�R ��	.��R��̖�R���!���O�,��-1R*a�P�RR���@����!������X�j�[� �"�b�B�	�'�l
�����e�i�B��	f����,�m!��SwĻ�M|N���K{�֧�5�M�Aㅐe����a&�q!�a���!�D�B�S6��
9��)��u.a\���빟#eSSk��< ��Z�7<� �O�ژaD�}V��Va;���O����Q�����9Q����U{��x���������0�sF	FЀ��N��p
�VL��2$c��u\PѨh�DKFh�h�wհ��cbS� ʥ�=���]�x����$�^P��W�:l�U��Zo�6�Umt����2&��e��E�ܖ����'UIw$Ō�) ���p��dE Zd�aiQ9�s�|�I���*G��:��ۄc��u�#�Q�T���ל��m�b�W�1���U�l�N�F��2߫��aUq�t߸�Z�Q��z>����79����:Aӟ����x��������#E6DH�K�GV`Vq~V���/Y��u��i5l��k<�K��Ⱦ,�^��?���v�7�J�ܗ��)��߽�$�p{��0�()�tÐ"*�,�@Z�Ԗ�m�𨍺�Hm���Cf��� ��$P���.��p�l�2_s��<	�èL�sT���A4����Kw����I���d]��1�'g��o4�ln�˱]��+��k،����j�e:>�s!�1��G�=�	��o�o-�/YJ3�o&W'��5�UV��;q1A�d�����γu�l���*t��z����tn(����"���1�G��!� L�k�H�T&�J�.�	I���H����%K豋V��3	�g^���|{r���տ�D6K~�K�/��`�-t��Ik�ثB�ln��H�0�cY.�%�;������nu��/���\�4�wR�e5"��.�}�hQ{��6=�+�ϢY��j���=�Qg�G�I���nq,?�f���|��Ғ���+���'�&X��.�`!��}���nSD����f�Ъ��C0SV�s�9��rZ�;�E��2��J���ɾOıy��G��)����W'�p0.{+d@�S�X����ӛn�Ȇyֆ&Ϗv�]#�_!9 (1]�����5��U�OAd���,�s�d�9㋿(4��q�:�_��Z�V-����j��~��K������ތyNL���Y36�:��hz��ёFu|�(��'�l��\���wY[�`s�6S��VI	���=��Gϗ]�෧hLj��|۞m�4l��Ҭ�,����5vo�<�����	{���=ES'F@�"_�v��	[�f���]��!V3�-���^�YP��t�6"ny��v���o^)i�q%&�7	l�l˫1��#�hu
��󜆘�EC̲^��.��_�M3�#bnS�'�٦?as]2����GVs�������dȰ     