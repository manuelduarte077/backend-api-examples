defmodule Shop.Repo.Migrations.CreateTableProducts do
  use Ecto.Migration

  def change do
    create table("products") do
      add :title, :string
      add :description, :string
      add :price, :integer, default: 0

      timestamps()
    end
  end
end
