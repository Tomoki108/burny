```mermaid
erDiagram
    projects ||--o{ sprints: "projectは0個以上のsprintを持つ"
    sprints ||--o| sprint_stats: "sprintは0個または1個のsprint_statを持つ"
    sprint_stats

```
