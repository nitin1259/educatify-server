CREATE DATABASE  `blog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE blog;


DROP TABLE IF EXISTS category;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE category_seq;

CREATE TABLE category (
  id bigint NOT NULL DEFAULT NEXTVAL ('category_seq'),
  parentId bigint DEFAULT NULL,
  title varchar(75) NOT NULL,
  metaTitle varchar(100) DEFAULT NULL,
  slug varchar(100) NOT NULL,
  content text,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_category_parent FOREIGN KEY (parentId) REFERENCES category (id)
)  ;

CREATE INDEX idx_category_parent ON category (parentId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES category WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS post;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE post_seq;

CREATE TABLE post (
  id bigint NOT NULL DEFAULT NEXTVAL ('post_seq'),
  authorId bigint NOT NULL,
  parentId bigint DEFAULT NULL,
  title varchar(75) NOT NULL,
  metaTitle varchar(100) DEFAULT NULL,
  slug varchar(100) NOT NULL,
  summary tinytext,
  published smallint NOT NULL DEFAULT '0',
  createdAt timestamp(0) NOT NULL,
  updatedAt timestamp(0) DEFAULT NULL,
  publishedAt timestamp(0) DEFAULT NULL,
  content text,
  PRIMARY KEY (id),
  CONSTRAINT uq_slug UNIQUE  (slug)
 ,
  CONSTRAINT fk_post_parent FOREIGN KEY (parentId) REFERENCES post (id),
  CONSTRAINT fk_post_user FOREIGN KEY (authorId) REFERENCES user (id)
)  ;

CREATE INDEX idx_post_user ON post (authorId);
CREATE INDEX idx_post_parent ON post (parentId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES post WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_category`
--

DROP TABLE IF EXISTS post_category;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE TABLE post_category (
  postId bigint NOT NULL,
  categoryId bigint NOT NULL,
  PRIMARY KEY (postId,categoryId)
  /*!80000 INVISIBLE */,
  CONSTRAINT fk_pc_category FOREIGN KEY (categoryId) REFERENCES category (id),
  CONSTRAINT fk_pc_post FOREIGN KEY (postId) REFERENCES post (id)
)  ;

CREATE INDEX idx_pc_category ON post_category (categoryId);
CREATE INDEX idx_pc_post ON post_category (postId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_category`
--

LOCK TABLES post_category WRITE;
/*!40000 ALTER TABLE `post_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_comment`
--

DROP TABLE IF EXISTS post_comment;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE post_comment_seq;

CREATE TABLE post_comment (
  id bigint NOT NULL DEFAULT NEXTVAL ('post_comment_seq'),
  postId bigint NOT NULL,
  parentId bigint DEFAULT NULL,
  title varchar(100) NOT NULL,
  published smallint NOT NULL DEFAULT '0',
  createdAt timestamp(0) NOT NULL,
  publishedAt timestamp(0) DEFAULT NULL,
  content text,
  PRIMARY KEY (id)
  /*!80000 INVISIBLE */
 ,
  CONSTRAINT fk_comment_parent FOREIGN KEY (parentId) REFERENCES post_comment (id),
  CONSTRAINT fk_comment_post FOREIGN KEY (postId) REFERENCES post (id)
)  ;

CREATE INDEX idx_comment_post ON post_comment (postId);
CREATE INDEX idx_comment_parent ON post_comment (parentId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_comment`
--

LOCK TABLES post_comment WRITE;
/*!40000 ALTER TABLE `post_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_meta`
--

DROP TABLE IF EXISTS post_meta;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE post_meta_seq;

CREATE TABLE post_meta (
  id bigint NOT NULL DEFAULT NEXTVAL ('post_meta_seq'),
  postId bigint NOT NULL,
  key varchar(50) NOT NULL,
  content text,
  PRIMARY KEY (id),
  CONSTRAINT uq_post_meta UNIQUE  (postId,key) /*!80000 INVISIBLE */
 ,
  CONSTRAINT fk_meta_post FOREIGN KEY (postId) REFERENCES post (id)
)  ;

CREATE INDEX idx_meta_post ON post_meta (postId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_meta`
--

LOCK TABLES post_meta WRITE;
/*!40000 ALTER TABLE `post_meta` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_meta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_tag`
--

DROP TABLE IF EXISTS post_tag;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE TABLE post_tag (
  postId bigint NOT NULL,
  tagId bigint NOT NULL,
  PRIMARY KEY (postId,tagId)
 ,
  CONSTRAINT fk_pt_post FOREIGN KEY (postId) REFERENCES post (id),
  CONSTRAINT fk_pt_tag FOREIGN KEY (tagId) REFERENCES tag (id)
)  ;

CREATE INDEX idx_pt_tag ON post_tag (tagId);
CREATE INDEX idx_pt_post ON post_tag (postId);
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_tag`
--

LOCK TABLES post_tag WRITE;
/*!40000 ALTER TABLE `post_tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `post_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS tag;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE tag_seq;

CREATE TABLE tag (
  id bigint NOT NULL DEFAULT NEXTVAL ('tag_seq'),
  title varchar(75) NOT NULL,
  metaTitle varchar(100) DEFAULT NULL,
  slug varchar(100) NOT NULL,
  content text,
  PRIMARY KEY (id)
)  ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES tag WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS user;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 character_set_client := utf8mb4 ;
CREATE SEQUENCE user_seq;

CREATE TABLE user (
  id bigint NOT NULL DEFAULT NEXTVAL ('user_seq'),
  firstName varchar(50) DEFAULT NULL,
  middleName varchar(50) DEFAULT NULL,
  lastName varchar(50) DEFAULT NULL,
  mobile varchar(15) DEFAULT NULL,
  email varchar(50) DEFAULT NULL,
  passwordHash varchar(32) NOT NULL,
  registeredAt timestamp(0) NOT NULL,
  lastLogin timestamp(0) DEFAULT NULL,
  intro tinytext,
  profile text,
  PRIMARY KEY (id),
  CONSTRAINT uq_mobile UNIQUE  (mobile),
  CONSTRAINT uq_emai UNIQUE  (email)
)  ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES user WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;