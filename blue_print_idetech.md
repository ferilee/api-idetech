---
title: "Blueprint Pengembangan Ekosistem Digital IdeTech"
subtitle: "Versi v9.0 - Final"
author:
  - "Gunanto"
  - "Feri Dwi Hermawan"
  - "Aan Triono"
date: "31 Maret 2026"
toc: false
---

```{=latex}
\begin{titlepage}
\thispagestyle{empty}
\pagecolor{idetechblack}
\color{white}
\centering

\vspace*{0.8cm}
{\Large \textcolor{idetechgold}{INDONESIA DIGITAL TEACHER}}\\[0.35cm]
{\rule{0.78\textwidth}{1.1pt}}\\[1.1cm]

\includegraphics[width=0.36\textwidth]{/home/pgun/dev/idetech/logo.png}\\[0.9cm]

{\Huge \textbf{Blueprint Pengembangan}}\\[0.15cm]
{\Huge \textbf{Ekosistem Digital IdeTech}}\\[0.45cm]
{\Large \textcolor{idetechgold}{Versi v9.0 - Final}}\\[0.5cm]
{\large Single Source of Truth}\\[1.0cm]

\noindent
\colorbox{idetechpanel}{
\parbox{0.9\textwidth}{
\vspace{0.35cm}
\centering
{\normalsize Dokumen acuan utama pengembangan backend, frontend, infrastruktur, multi-tenant, AI, dan backup}\\
\vspace{0.35cm}
}}
\\[1.0cm]

\renewcommand{\arraystretch}{1.45}
\arrayrulecolor{idetechgold}
\begin{tabular}{|p{4.1cm}|p{8.9cm}|}
\hline
\rowcolor{idetechpanel}\textcolor{idetechgold}{\textbf{Tanggal}} & 31 Maret 2026 \\
\hline
\rowcolor{idetechpanel}\textcolor{idetechgold}{\textbf{Waktu Pembuatan}} & 31 Maret 2026 10:13:45 WIB \\
\hline
\rowcolor{idetechpanel}\textcolor{idetechgold}{\textbf{Tim Penyusun}} & Gunanto, Feri Dwi Hermawan, Aan Triono \\
\hline
\rowcolor{idetechpanel}\textcolor{idetechgold}{\textbf{Produk}} & IdeTech \\
\hline
\rowcolor{idetechpanel}\textcolor{idetechgold}{\textbf{Ruang Lingkup}} & Backend, Frontend, Infrastruktur, Multi-tenant, AI, Backup \\
\hline
\end{tabular}

\vfill
{\textcolor{idetechgold}{\rule{0.72\textwidth}{0.8pt}}}\\[0.3cm]
{\large \textbf{Dokumen Internal}}\\[0.15cm]
{\normalsize Untuk perencanaan implementasi sistem dan deployment}
\pagecolor{white}
\color{black}
\end{titlepage}

\tableofcontents
\newpage
```

## 1. Ringkasan Eksekutif

IdeTech adalah ekosistem pembelajaran digital multi-tenant untuk sekolah yang menggabungkan manajemen pembelajaran, gamifikasi, analitik visual, kolaborasi, dan otomasi konten berbasis AI. Platform ini dirancang modular agar dapat dikembangkan bertahap, di-deploy efisien pada VPS, dan disesuaikan untuk berbagai jenjang pendidikan.

Dokumen ini menjadi acuan utama untuk:

- pengembangan backend, frontend, dan infrastruktur;
- penyusunan skema data dan kontrak integrasi;
- deployment, backup, keamanan, dan scaling;
- menjaga konsistensi implementasi selama proses coding.

## 2. Visi Produk

IdeTech dibangun sebagai platform pembelajaran masa depan dengan prinsip:

- **Adaptif:** pengalaman belajar menyesuaikan peran, jenjang, dan performa siswa.
- **Visual:** data belajar mudah dibaca melalui dashboard analitik yang jelas.
- **Menyenangkan:** aktivitas belajar dikemas dalam bentuk quest, XP, badge, dan peta petualangan.
- **Kolaboratif:** guru dan siswa dapat berbagi ide, materi, dan ruang kerja bersama.
- **Aman dan terukur:** multi-tenant, terisolasi, mudah di-backup, dan siap ditingkatkan kapasitasnya.

## 3. Sasaran Utama Sistem

- Menyediakan portal digital per sekolah berbasis subdomain.
- Mendukung login cerdas berdasarkan peran: `guru`, `siswa`, `orang_tua`, dan `admin tenant`.
- Memungkinkan guru mengunggah materi, membuat aktivitas, dan menghasilkan kuis otomatis dengan AI.
- Menyediakan jalur belajar personal berbasis performa siswa.
- Menampilkan analitik belajar secara visual untuk intervensi dini.
- Menyediakan ruang kolaborasi digital dan marketplace materi guru.

## 4. Fitur Inti Produk

### 4.1 Login Cerdas dan Adaptif

- Sistem login menyesuaikan identitas tenant dari subdomain.
- Setelah autentikasi, pengguna diarahkan ke dashboard sesuai peran.
- Tampilan dan navigasi dapat dibedakan untuk SD, SMP, SMA, dan peran pengguna.

### 4.2 Personalisasi Visual Sesuai Jenjang

- Tema warna, ikon, logo, ilustrasi, dan kepadatan UI dikonfigurasi per tenant.
- SD menggunakan visual lebih ceria dan sederhana.
- SMP/SMA menggunakan tampilan lebih ringkas, fokus, dan informatif.

### 4.3 IdeStudio

- Modul guru untuk mengelola materi, unggah file, dan membuat konten pembelajaran.
- Mendukung unggah PDF, video, dokumen, dan aset pendukung.

### 4.4 IdeGen

- Modul AI untuk mengubah materi menjadi kuis, ringkasan, atau latihan.
- Integrasi awal menggunakan Gemini API.
- Output utama berupa soal pilihan ganda dalam format JSON terstruktur.

### 4.5 IdeQuest

- Materi dan tugas ditampilkan sebagai quest belajar.
- Setiap quest dapat memberikan XP, badge, dan progres.
- Quest dapat memiliki prasyarat untuk membentuk jalur belajar.

### 4.6 Jalur Belajar Personal

- Sistem mendeteksi performa siswa dari hasil kuis dan aktivitas.
- Jika skor di bawah ambang tertentu, sistem membuka materi remedial.
- Jika skor baik, sistem dapat membuka jalur pengayaan.

### 4.7 Co-Lab

- Ruang kolaborasi digital untuk diskusi, berbagi file, dan tugas kelompok.
- Dapat berkembang ke fitur sinkron real-time berbasis WebSocket.

### 4.8 Bank Ide

- Marketplace materi guru untuk menyimpan, berbagi, dan mengkurasi aset pembelajaran.
- Mendukung penilaian, kategori, dan pencarian.

### 4.9 Radar Pintar

- Dashboard analitik visual untuk guru dan admin sekolah.
- Fokus pada performa siswa, distribusi nilai, area lemah, dan potensi intervensi dini.

## 5. Stack Teknologi

| Layer | Teknologi | Peran |
|---|---|---|
| Backend | Go (Golang) + Chi Router | API cepat, ringan, dan mudah dipelihara |
| Frontend | Next.js 14+ App Router | SSR/ISR, routing modern, dan cocok untuk dashboard |
| Styling | Tailwind CSS + shadcn/ui | UI modular, cepat dibangun, dan konsisten |
| Database | PostgreSQL 16 | Sumber data relasional utama |
| Cache | Redis | sesi, leaderboard, flag personalisasi, dan cache |
| Object Storage | MinIO | penyimpanan file materi dan aset |
| Reverse Proxy | Nginx Proxy Manager | SSL, domain, dan reverse proxy dengan UI |
| Containerization | Docker + Docker Compose | konsistensi runtime lokal dan server |
| Registry | GitHub Container Registry (GHCR) | distribusi image |
| AI Engine | Gemini API | generasi kuis dan konten otomatis |
| Backup | pg_dump + GDrive API + Telegram Bot | backup otomatis dan redundansi |

## 6. Arsitektur Sistem Tingkat Tinggi

Arsitektur IdeTech dibangun dengan pendekatan service modular:

1. **Frontend Web**
   Next.js melayani UI tenant, dashboard peran, dan halaman publik.

2. **API Backend**
   Go bertanggung jawab atas autentikasi, manajemen tenant, pembelajaran, gamifikasi, file, dan analitik.

3. **Data Layer**
   PostgreSQL sebagai penyimpanan utama; Redis sebagai state cepat dan cache.

4. **Storage Layer**
   MinIO menyimpan file PDF, video, gambar, avatar, dan lampiran kolaborasi.

5. **AI Integration Layer**
   Backend menjadi penghubung aman antara data tenant dan Gemini API.

6. **Ops Layer**
   Docker Compose, GHCR, Nginx Proxy Manager, cron backup, dan monitoring dasar.

## 7. Arsitektur Multi-Tenancy

IdeTech menggunakan pendekatan **single database, shared schema, tenant-aware rows**.

### Prinsip utama

- Setiap entitas bisnis wajib memiliki `tenant_id`, kecuali tabel sistem tertentu.
- Semua query aplikasi wajib difilter berdasarkan `tenant_id`.
- Tenant diidentifikasi dari subdomain, misalnya `smpn1punggur.idetech.id`.
- Frontend memuat konfigurasi visual tenant dari API saat bootstrap.

### Keunggulan pendekatan ini

- lebih hemat sumber daya server;
- lebih cepat diimplementasikan;
- memudahkan onboarding sekolah baru;
- tetap aman selama disiplin filtering dan authorization diterapkan.

### Risiko dan mitigasi

- Risiko kebocoran data antar tenant jika query tidak difilter.
- Mitigasi melalui middleware tenant, repository pattern, test isolasi data, dan audit query.

## 8. Prinsip Backend

### Struktur modul utama backend

- `auth`
- `tenant`
- `user`
- `quest`
- `material`
- `quiz`
- `analytics`
- `collab`
- `storage`
- `backup`
- `notification`

### Standar implementasi

- Gunakan `Chi Router` untuk routing dan middleware.
- Gunakan struktur berlapis: `handler -> service -> repository`.
- Semua endpoint private memakai JWT.
- Tenant context diekstrak lebih awal di middleware.
- Gunakan migrasi database versi terkontrol.
- Semua operasi penting menghasilkan audit log.

## 9. Prinsip Frontend

### Arah implementasi

- Next.js App Router untuk struktur route modern.
- Server Components untuk data statis/umum, Client Components untuk interaksi tinggi.
- shadcn/ui dipakai sebagai basis komponen, lalu dikustomisasi sesuai identitas IdeTech.

### Zona utama frontend

- halaman login tenant-aware;
- dashboard guru;
- dashboard siswa;
- dashboard orang tua;
- dashboard admin tenant;
- halaman quest map;
- halaman analitik;
- halaman Co-Lab;
- halaman Bank Ide;
- halaman pengaturan tenant.

### Prinsip UI

- responsif di desktop dan mobile;
- visual per jenjang dapat dikustom dari `tenant.config`;
- navigasi sederhana untuk siswa, lebih kaya fitur untuk guru/admin;
- dark mode opsional, tetapi bukan prioritas awal dibanding tema tenant.

## 10. Model Domain Inti

Entitas utama yang akan menjadi fondasi sistem:

- `tenants`
- `users`
- `roles`
- `materials`
- `quests`
- `quizzes`
- `quiz_questions`
- `quiz_attempts`
- `student_progress`
- `leaderboards`
- `collab_rooms`
- `collab_messages`
- `idea_market_items`
- `audit_logs`
- `backup_logs`

## 11. Skema Database Inti

Contoh skema awal yang dapat dijadikan baseline:

```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(50) NOT NULL UNIQUE,
    domain VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    config JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255),
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL,
    profile_data JSONB NOT NULL DEFAULT '{}'::jsonb,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, username),
    UNIQUE (tenant_id, email)
);

CREATE TABLE materials (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    created_by UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    file_url TEXT NOT NULL,
    file_type VARCHAR(50),
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE quests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    xp_reward INT NOT NULL DEFAULT 0,
    material_id UUID REFERENCES materials(id),
    parent_quest_id UUID REFERENCES quests(id),
    difficulty_level INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE quizzes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    quest_id UUID REFERENCES quests(id) ON DELETE SET NULL,
    material_id UUID REFERENCES materials(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    source VARCHAR(30) NOT NULL DEFAULT 'manual',
    config JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE quiz_questions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    question_text TEXT NOT NULL,
    question_type VARCHAR(30) NOT NULL DEFAULT 'multiple_choice',
    options JSONB NOT NULL,
    answer_key JSONB NOT NULL,
    explanation TEXT,
    position INT NOT NULL
);

CREATE TABLE quiz_attempts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    quiz_id UUID NOT NULL REFERENCES quizzes(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    score NUMERIC(5,2) NOT NULL DEFAULT 0,
    submitted_answers JSONB NOT NULL DEFAULT '{}'::jsonb,
    result_summary JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE student_progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    quest_id UUID NOT NULL REFERENCES quests(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'locked',
    xp_earned INT NOT NULL DEFAULT 0,
    completed_at TIMESTAMP,
    UNIQUE (tenant_id, user_id, quest_id)
);
```

### Catatan desain data

- Tambahkan index pada `tenant_id` di seluruh tabel tenant-aware.
- Pertimbangkan index gabungan seperti `(tenant_id, role)`, `(tenant_id, created_at)`, dan `(tenant_id, user_id)`.
- Simpan konfigurasi fleksibel di `JSONB`, tetapi data bisnis penting tetap dalam kolom terstruktur.

## 12. Alur Sistem Utama

### 12.1 Alur Login dan Resolusi Tenant

1. User membuka subdomain sekolah.
2. Frontend mengirim request bootstrap ke backend.
3. Backend mengekstrak slug tenant dari host.
4. Backend memuat data tenant dan konfigurasi tema.
5. Frontend merender UI sesuai tenant.
6. Setelah login berhasil, user diarahkan sesuai peran.

### 12.2 Alur IdeStudio dan IdeGen

1. Guru unggah file ke backend.
2. Backend menyimpan file ke MinIO.
3. Metadata file disimpan di PostgreSQL.
4. Guru memilih aksi "generate quiz".
5. Backend mengambil konten file yang relevan.
6. Backend mengirim prompt terstruktur ke Gemini API.
7. Respons AI divalidasi dan dinormalisasi.
8. Kuis disimpan ke `quizzes` dan `quiz_questions`.

### 12.3 Alur Quest dan Personalisasi

1. Siswa membuka peta belajar.
2. Backend mengembalikan daftar quest berdasarkan progres.
3. Saat kuis selesai, sistem menghitung skor.
4. Jika skor `< 70`, Redis menandai status remedial.
5. Frontend membuka node materi bantuan atau jalur remedial.
6. Jika skor memenuhi target, jalur lanjutan dibuka.

### 12.4 Alur Radar Pintar

1. Guru memilih kelas atau mata pelajaran.
2. Backend menjalankan query agregasi PostgreSQL.
3. Data diringkas menjadi metrik visual.
4. Frontend menampilkan radar chart, bar chart, dan indikator risiko.
5. Guru menerima rekomendasi intervensi berbasis data.

### 12.5 Alur Co-Lab

1. Guru membuat ruang kolaborasi.
2. Siswa masuk ke room sesuai kelas atau kelompok.
3. Sistem mendukung posting, komentar, dan lampiran file.
4. Fase lanjut menggunakan WebSocket untuk sinkronisasi real-time.

### 12.6 Alur Bank Ide

1. Guru mengunggah materi ke marketplace internal tenant.
2. Materi dikategorikan, diberi tag, dan dapat dinilai.
3. Pengguna lain dapat melihat, menggunakan, atau mengadaptasi materi.

## 13. Redis Strategy

Redis digunakan untuk kebutuhan cepat dan sementara:

- session cache;
- refresh token blacklist atau session metadata;
- leaderboard per tenant;
- flag `needs_remedial`;
- cache dashboard analitik ringan;
- rate limit sederhana;
- state real-time untuk Co-Lab dan notifikasi.

Contoh key naming:

```text
tenant:{tenant_id}:leaderboard
tenant:{tenant_id}:user:{user_id}:needs_remedial
tenant:{tenant_id}:quiz:{quiz_id}:live
tenant:{tenant_id}:theme
```

## 14. MinIO Strategy

### Jenis file yang disimpan

- PDF materi;
- video pembelajaran;
- thumbnail;
- avatar;
- file tugas;
- lampiran Co-Lab.

### Konvensi bucket/object

- bucket per lingkungan: `idetech-dev`, `idetech-prod`;
- object path per tenant:
  `tenant/{tenant_slug}/materials/...`
  `tenant/{tenant_slug}/avatars/...`
  `tenant/{tenant_slug}/collab/...`

### Prinsip akses

- file private diakses via signed URL;
- metadata file tetap dicatat di PostgreSQL;
- validasi MIME type dan ukuran file wajib diterapkan.

## 15. Kontrak API Awal

Contoh endpoint minimum viable platform:

- `POST /api/v1/auth/login`
- `GET /api/v1/auth/me`
- `POST /api/v1/tenants`
- `GET /api/v1/tenant/bootstrap`
- `GET /api/v1/users`
- `POST /api/v1/materials/upload`
- `POST /api/v1/materials/:id/generate-quiz`
- `GET /api/v1/quests`
- `POST /api/v1/quizzes/:id/submit`
- `GET /api/v1/analytics/radar`
- `GET /api/v1/leaderboard`
- `POST /api/v1/collab/rooms`
- `GET /api/v1/idea-market`

## 16. Keamanan dan Governance

### Minimum security baseline

- JWT access token dengan expiry pendek;
- refresh token aman dan dapat dicabut;
- password di-hash dengan Argon2id atau bcrypt biaya memadai;
- validasi role dan permission di setiap endpoint;
- sanitasi input dan validasi payload;
- rate limiting untuk login dan endpoint sensitif;
- signed URL untuk file private;
- audit log untuk login, upload, generate quiz, dan backup.

### Multi-tenant security rules

- jangan pernah percaya `tenant_id` dari body request tanpa validasi context;
- ambil tenant dari host atau token;
- batasi query repository berdasarkan tenant context;
- tambahkan test integrasi untuk memastikan isolasi tenant.

## 17. Observability dan Monitoring

Target minimum fase awal:

- structured logging pada backend;
- health check endpoint;
- log rotasi container;
- status backup tercatat;
- notifikasi kegagalan backup ke Telegram;
- metrik dasar CPU, memory, disk, dan container uptime.

Tool observability dapat ditingkatkan bertahap ke:

- Prometheus;
- Grafana;
- Loki atau stack log lain.

## 18. DevOps dan Deployment

### Container utama

- `frontend`
- `backend`
- `postgres`
- `redis`
- `minio`
- `nginx-proxy-manager`
- `backup-worker`

### Prinsip deployment

- semua service berjalan via Docker Compose;
- image backend/frontend dibangun via GitHub Actions;
- hasil build dipush ke GHCR;
- server menarik image terbaru saat rilis;
- environment dipisahkan minimal `development` dan `production`.

### Struktur environment variable utama

- `APP_ENV`
- `APP_URL`
- `JWT_SECRET`
- `DATABASE_URL`
- `REDIS_URL`
- `MINIO_ENDPOINT`
- `MINIO_ACCESS_KEY`
- `MINIO_SECRET_KEY`
- `MINIO_BUCKET`
- `GEMINI_API_KEY`
- `GDRIVE_CREDENTIALS_JSON`
- `TELEGRAM_BOT_TOKEN`
- `TELEGRAM_CHAT_ID`

## 19. Roadmap Implementasi 7 Minggu

### Minggu 1: Infrastructure dan DevOps

- siapkan VPS berbasis Ubuntu atau setara;
- pasang Docker dan Docker Compose;
- pasang Nginx Proxy Manager;
- siapkan repo dan workflow GitHub Actions;
- push image ke GHCR.

### Minggu 2: Core Backend

- implementasi auth JWT;
- middleware tenant resolver;
- CRUD tenant;
- koneksi PostgreSQL dan Redis;
- migrasi database awal.

### Minggu 3: Frontend Foundation

- inisialisasi Next.js + Tailwind + shadcn/ui;
- halaman login multi-tenant;
- layout dashboard per peran;
- theme loader dari `tenant.config`.

### Minggu 4: Content Engine

- upload file ke MinIO;
- halaman daftar materi;
- integrasi IdeGen dengan Gemini API;
- validasi hasil quiz generation.

### Minggu 5: Gamification Engine

- quest map siswa;
- XP, badge, dan progress tracking;
- leaderboard Redis;
- remedial dan pengayaan adaptif.

### Minggu 6: Radar Pintar dan Co-Lab

- dashboard analitik guru;
- visualisasi radar dan batang;
- room kolaborasi;
- WebSocket dasar bila waktu mencukupi.

### Minggu 7: Hardening dan Backup

- backup otomatis database dan konfigurasi penting;
- unggah ke Google Drive;
- kirim ke Telegram;
- audit keamanan dasar;
- finalisasi deployment production.

## 20. Strategi Backup

### Tujuan

Menjamin data tenant, konfigurasi sistem, dan artefak penting dapat dipulihkan dengan cepat jika terjadi kegagalan server, kesalahan operasional, atau kerusakan data.

### Ruang lingkup backup

- dump PostgreSQL;
- metadata aplikasi penting;
- konfigurasi Nginx Proxy Manager;
- file environment yang aman untuk diarsipkan;
- opsional: snapshot metadata MinIO atau sinkronisasi bucket terpisah.

### Jadwal

- backup harian pada pukul `02:00`;
- backup mingguan dengan retensi lebih panjang;
- notifikasi hasil backup dikirim setelah job selesai.

### Mekanisme

1. Worker menjalankan `pg_dump`.
2. File backup dan konfigurasi dibungkus ke `.tar.gz`.
3. Arsip dienkripsi.
4. Arsip diunggah ke Google Drive.
5. Salinan dikirim ke grup admin privat via Telegram.
6. Status dicatat di `backup_logs`.

### Retensi yang disarankan

- harian: 7 sampai 14 hari;
- mingguan: 8 minggu;
- bulanan: 6 sampai 12 bulan.

### Catatan penting

- Jika ukuran backup membesar, Telegram bisa dipakai untuk notifikasi dan checksum, bukan file penuh.
- Untuk aset MinIO yang besar, gunakan strategi sinkronisasi bucket terjadwal, bukan mengompres semua file setiap hari.

## 21. Disaster Recovery

Target minimum:

- restore database dari backup terenkripsi;
- restore konfigurasi proxy dan environment;
- verifikasi aplikasi dapat boot;
- validasi domain/subdomain tenant aktif kembali;
- uji restore dilakukan berkala, bukan hanya backup.

RTO/RPO awal yang realistis:

- **RTO target:** 2 sampai 4 jam.
- **RPO target:** maksimum kehilangan data 24 jam pada fase awal.

## 22. Prioritas MVP

Fitur yang wajib ada pada rilis awal:

- multi-tenant berbasis subdomain;
- login dan role-based dashboard;
- CRUD tenant, user, material, dan quest;
- upload file ke MinIO;
- generate quiz via AI;
- quiz submission dan score;
- quest map dasar;
- radar analitik dasar;
- backup otomatis minimum.

Fitur yang bisa menyusul setelah MVP:

- marketplace materi lintas tenant;
- kolaborasi real-time penuh;
- parent portal lanjutan;
- mobile app native;
- rekomendasi AI yang lebih mendalam.

## 23. Standar Kualitas Engineering

### Backend

- unit test untuk service penting;
- integration test untuk auth, tenant filter, dan quiz submission;
- migration tested;
- lint dan format otomatis.

### Frontend

- component consistency;
- loading, empty, dan error state jelas;
- aksesibilitas dasar terpenuhi;
- responsif untuk perangkat sekolah umum.

### Ops

- `.env.example` lengkap;
- health check tersedia;
- script deploy terdokumentasi;
- prosedur restore backup terdokumentasi.

## 24. Struktur Folder yang Disarankan

```text
idetech/
|-- backend/
|   |-- cmd/
|   |-- internal/
|   |   |-- auth/
|   |   |-- tenant/
|   |   |-- user/
|   |   |-- material/
|   |   |-- quest/
|   |   |-- quiz/
|   |   |-- analytics/
|   |   |-- collab/
|   |   |-- backup/
|   |   `-- platform/
|   `-- migrations/
|-- frontend/
|   |-- app/
|   |-- components/
|   |-- lib/
|   `-- styles/
|-- deploy/
|   |-- docker/
|   |-- nginx/
|   |-- scripts/
|   `-- backup/
|-- docs/
`-- blue_print_idetech.md
```

## 25. Keputusan Arsitektur yang Dikunci

Keputusan berikut dianggap final kecuali ada alasan teknis kuat untuk diubah:

- backend utama menggunakan Go + Chi;
- frontend utama menggunakan Next.js App Router;
- database utama menggunakan PostgreSQL;
- cache dan transient state menggunakan Redis;
- object storage menggunakan MinIO;
- multi-tenant menggunakan shared schema dengan `tenant_id`;
- deploy utama menggunakan Docker Compose;
- reverse proxy menggunakan Nginx Proxy Manager;
- AI generation awal menggunakan Gemini API.

## 26. Catatan Penyempurnaan dari Infografis

Poin-poin infografis yang menjadi penegas desain produk:

- **Login Cerdas & Adaptif:** sistem harus sadar konteks tenant, peran, dan jenjang.
- **Personalisasi Visual:** UI SD dan SMA tidak boleh terasa identik.
- **IdeStudio:** guru harus bisa membuat konten tanpa bergantung pada skill teknis tinggi.
- **IdeQuest:** pembelajaran divisualkan sebagai peta petualangan, bukan sekadar daftar tugas.
- **Jalur Belajar Personal:** remedial dan pengayaan terbuka otomatis berdasarkan performa.
- **Co-Lab:** siswa dan guru punya ruang ide bersama, bukan hanya forum pasif.
- **Bank Ide:** materi harus bisa dibagikan, dikurasi, dan dimonetisasi internal bila dibutuhkan.
- **Radar Pintar:** analitik harus mengarah ke tindakan, bukan hanya grafik.

## 27. Urutan Implementasi yang Disarankan

Untuk menjaga fokus delivery, urutan pengerjaan yang paling aman:

1. bangun backend core dan skema database;
2. bangun auth, tenant resolver, dan bootstrap tenant;
3. bangun frontend login dan dashboard dasar;
4. sambungkan upload material ke MinIO;
5. integrasikan AI quiz generation;
6. bangun quest map dan progress siswa;
7. bangun dashboard analitik guru;
8. aktifkan Co-Lab dan fitur marketplace;
9. finalkan backup, hardening, dan deployment.

## 28. Penutup

Blueprint ini dirancang agar IdeTech dapat dikerjakan secara modular namun tetap terarah. Tahapan awal difokuskan pada fondasi yang kuat: auth, tenant isolation, materi, kuis, dan dashboard. Setelah itu, diferensiasi produk dibangun melalui gamifikasi, personalisasi visual, AI assistance, Co-Lab, dan Radar Pintar.

Jika terjadi konflik keputusan selama coding, dokumen ini menjadi rujukan utama sampai ada revisi versi baru yang disepakati.
