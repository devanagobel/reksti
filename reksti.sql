-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Apr 25, 2018 at 02:03 PM
-- Server version: 10.1.16-MariaDB
-- PHP Version: 7.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `reksti`
--

-- --------------------------------------------------------

--
-- Table structure for table `class`
--

CREATE TABLE `class` (
  `class_index` varchar(10) NOT NULL,
  `class_name` varchar(50) NOT NULL,
  `course_index` varchar(7) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `class`
--

INSERT INTO `class` (`class_index`, `class_name`, `course_index`) VALUES
('EL4233-01', 'Fuzzy Sets', 'EL4233'),
('EL4233-02', 'Fuzzy Relation', 'EL4233'),
('EL4233-03', 'Fuzzy Rules', 'EL4233'),
('EL4233-04', 'Fuzzy Reasoning', 'EL4233'),
('EL4233-05', 'Fuzzy Logic Control', 'EL4233'),
('EL4233-06', 'Machine Learning Introduction', 'EL4233'),
('EL4233-07', 'Machine Learning Linear Regression', 'EL4233'),
('EL4233-08', 'Machine Learning Polynomial Curve Fitting', 'EL4233'),
('EL4233-09', 'Machine Learning Training Linear Regression', 'EL4233'),
('EL4233-10', 'Machine Learning Application', 'EL4233');

-- --------------------------------------------------------

--
-- Table structure for table `course`
--

CREATE TABLE `course` (
  `course_index` varchar(7) NOT NULL,
  `course_name` varchar(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `course`
--

INSERT INTO `course` (`course_index`, `course_name`) VALUES
('EL4233', 'Dasar Sistem dan Kendali Cerdas'),
('II3220', 'Arsitektur Enterprise'),
('II3230', 'Keamanan Informasi'),
('II3240', 'Rekayasa Sistem dan Teknologi Informasi');

-- --------------------------------------------------------

--
-- Table structure for table `student`
--

CREATE TABLE `student` (
  `student_nim` varchar(10) NOT NULL,
  `student_name` varchar(35) NOT NULL,
  `student_faculty` varchar(35) NOT NULL,
  `student_major` varchar(35) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `student`
--

INSERT INTO `student` (`student_nim`, `student_name`, `student_faculty`, `student_major`) VALUES
('13215002', 'Yosua Hauw', 'FTI', 'Teknik Industri'),
('13515001', 'Gisela Supardi', 'STEI', 'Teknik Informatika'),
('18215004', 'Teo Wijayarto', 'STEI', 'Sistem Teknologi dan Informasi'),
('18215028', 'Devana Gobel', 'STEI', 'Sistem Teknologi dan Informasi'),
('18215047', 'Ikhsan Widi Adiyatma', 'STEI', 'Sistem Teknologi dan Informasi');

-- --------------------------------------------------------

--
-- Table structure for table `teacher`
--

CREATE TABLE `teacher` (
  `teacher_id` int(11) NOT NULL,
  `teacher_nip` varchar(10) NOT NULL,
  `teacher_name` varchar(35) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `class`
--
ALTER TABLE `class`
  ADD PRIMARY KEY (`class_index`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`course_index`);

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`student_nim`);

--
-- Indexes for table `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`teacher_id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
