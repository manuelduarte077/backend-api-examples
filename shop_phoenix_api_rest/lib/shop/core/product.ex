defmodule Shop.Core.Product do
  @moduledoc false

  use Ecto.Schema

  schema "products" do
    field(:title, :string)
    field(:description, :string)
    field(:price, :integer, default: 0)

    timestamps()
  end

  def permission(:public), do: [:id, :title, :description, :price, :inserted_at, :updated_at]

  def changeset(model, params) do
    model
    |> Ecto.Changeset.cast(params, [:title, :description, :price])
    |> Ecto.Changeset.validate_required([:title, :description])
  end
end
