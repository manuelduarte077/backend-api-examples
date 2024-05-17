defmodule Shop.Core.Product.Api do
  @moduledoc false
  alias Shop.Repo
  alias Shop.Core.Product

  import Ecto.Query

  ## all, get, get_by, insert, update, delete
  def all do
    {:ok, Repo.all(Product)}
  end

  def get(id) do
    Repo.get(Product, id)
    |> case do
      nil -> {:error, "Product not found"}
      product -> {:ok, product}
    end
  end

  def get_by(where) do
    from(Product)
    |> where(^where)
    |> Repo.one()
  end

  def delete(id) do
    case get(id) do
      {:ok, product} -> Repo.delete(product)
      :error -> {:error, "Product not found"}
    end
  end

  def insert(params) do
    %Product{}
    |> Product.changeset(params)
    |> Repo.insert()
  end

  def update(model, params) do
    model
    |> Product.changeset(params)
    |> Repo.update()
  end

  def json!(model, permission) do
    Map.take(model, Product.permission(permission))
  end
end
