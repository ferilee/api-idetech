import { TenantBootstrapForm } from "@/components/tenant-bootstrap-form";

const sections = [
  {
    title: "API Go + Chi",
    body: "Fondasi backend memuat route health check, tenant bootstrap, dan titik masuk untuk auth serta modul tenant-aware berikutnya."
  },
  {
    title: "UI Tenant-Aware",
    body: "Frontend Next menyiapkan pengalaman login dan dashboard yang bisa berubah sesuai identitas tenant dan jenjang."
  },
  {
    title: "Data Layer",
    body: "Migrasi awal PostgreSQL menurunkan entitas inti dari blueprint: tenant, user, material, quest, quiz, dan progress."
  }
];

export default function HomePage() {
  return (
    <main>
      <div className="shell">
        <section className="hero">
          <div className="panel hero-copy">
            <span className="eyebrow">Blueprint to Product</span>
            <h1>IdeTech mulai dari fondasi tenant-aware yang rapi.</h1>
            <p>
              Tahap awal ini menurunkan blueprint menjadi monorepo yang siap dikembangkan:
              backend Go, frontend Next, migrasi PostgreSQL, dan alur bootstrap tenant untuk
              konteks sekolah.
            </p>

            <div className="grid metrics">
              <div className="metric">
                <strong>7</strong>
                entitas data inti sudah diturunkan ke migrasi awal.
              </div>
              <div className="metric">
                <strong>1</strong>
                endpoint tenant bootstrap siap dipakai frontend.
              </div>
              <div className="metric">
                <strong>3</strong>
                layer dasar sudah ada: backend, frontend, deploy.
              </div>
              <div className="metric">
                <strong>MVP</strong>
                diarahkan ke auth, tenant resolver, dan dashboard dasar.
              </div>
            </div>
          </div>

          <TenantBootstrapForm />
        </section>

        <section className="panel sections">
          <h2>Milestone Bootstrap</h2>
          <div className="sections-grid">
            {sections.map((section) => (
              <article className="section-card" key={section.title}>
                <h3>{section.title}</h3>
                <p>{section.body}</p>
              </article>
            ))}
          </div>
        </section>
      </div>
    </main>
  );
}
